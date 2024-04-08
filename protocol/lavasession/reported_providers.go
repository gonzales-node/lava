package lavasession

import (
	"sync"
	"time"

	metrics "github.com/lavanet/lava/protocol/metrics"
	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
)

const (
	ReconnectCandidateTime = 2 * time.Minute
	debugReportedProviders = false
)

type ReportedProviders struct {
	addedToPurgeAndReport map[string]*ReportedProviderEntry // list of purged providers to report for QoS unavailability. (easier to search maps.)
	lock                  sync.RWMutex
	reporter              metrics.Reporter
}

type ReportedProviderEntry struct {
	Disconnections uint64
	Errors         uint64
	addedTime      time.Time
	reconnectCB    func() error
}

func (rp *ReportedProviders) Reset() {
	rp.lock.Lock()
	defer rp.lock.Unlock()
	if debugReportedProviders {
		utils.LavaFormatDebug("[debugReportedProviders] Reset called")
	}
	rp.addedToPurgeAndReport = make(map[string]*ReportedProviderEntry, 0)
}

func (rp *ReportedProviders) GetReportedProviders() []*pairingtypes.ReportedProvider {
	rp.lock.RLock()
	defer rp.lock.RUnlock()
	reportedProviders := make([]*pairingtypes.ReportedProvider, 0, len(rp.addedToPurgeAndReport))
	for provider, reportedProviderEntry := range rp.addedToPurgeAndReport {
		reportedProvider := pairingtypes.ReportedProvider{
			Address:        provider,
			Disconnections: reportedProviderEntry.Disconnections,
			Errors:         reportedProviderEntry.Errors,
			TimestampS:     reportedProviderEntry.addedTime.Unix(),
		}
		reportedProviders = append(reportedProviders, &reportedProvider)
	}
	return reportedProviders
}

func (rp *ReportedProviders) ReportProvider(address string, errors uint64, disconnections uint64, reconnectCB func() error) {
	rp.lock.Lock()
	defer rp.lock.Unlock()
	if _, ok := rp.addedToPurgeAndReport[address]; !ok { // add if it doesn't exist already
		utils.LavaFormatInfo("Reporting Provider for unresponsiveness", utils.Attribute{Key: "Provider address", Value: address})
		rp.addedToPurgeAndReport[address] = &ReportedProviderEntry{}
		rp.addedToPurgeAndReport[address].addedTime = time.Now()
	}
	rp.addedToPurgeAndReport[address].Disconnections += disconnections
	rp.addedToPurgeAndReport[address].Errors += errors
	if reconnectCB != nil {
		rp.addedToPurgeAndReport[address].reconnectCB = reconnectCB
	}
	if debugReportedProviders {
		utils.LavaFormatDebug("[debugReportedProviders] adding provider to reported providers", utils.LogAttr("rp.addedToPurgeAndReport", rp.addedToPurgeAndReport))
	}
}

// will be called after a disconnected provider got a valid connection
func (rp *ReportedProviders) RemoveReport(address string) {
	rp.lock.Lock()
	defer rp.lock.Unlock()
	if debugReportedProviders {
		utils.LavaFormatDebug("[debugReportedProviders] Removing Report", utils.LogAttr("address", address))
	}
	delete(rp.addedToPurgeAndReport, address)
}

func (rp *ReportedProviders) IsReported(address string) bool {
	rp.lock.RLock()
	defer rp.lock.RUnlock()
	_, ok := rp.addedToPurgeAndReport[address]
	return ok
}

type reconnectCandidate struct {
	address     string
	reconnectCB func() error
}

func (rp *ReportedProviders) ReconnectCandidates() []reconnectCandidate {
	rp.lock.RLock()
	defer rp.lock.RUnlock()
	candidates := []reconnectCandidate{}
	if debugReportedProviders {
		utils.LavaFormatDebug("[debugReportedProviders] Reconnect candidates", utils.LogAttr("candidate list", rp.addedToPurgeAndReport))
	}
	for address, entry := range rp.addedToPurgeAndReport {
		// only reconnect providers that didn't have consecutive errors
		if entry.Errors == 0 && time.Since(entry.addedTime) > ReconnectCandidateTime {
			candidate := reconnectCandidate{
				address:     address,
				reconnectCB: entry.reconnectCB,
			}
			candidates = append(candidates, candidate)
		}
	}
	return candidates
}

func (rp *ReportedProviders) ReconnectProviders() {
	candidates := rp.ReconnectCandidates()
	for _, candidate := range candidates {
		if candidate.reconnectCB != nil {
			if debugReportedProviders {
				utils.LavaFormatDebug("[debugReportedProviders] Trying to reconnect candidate", utils.LogAttr("candidate", candidate.address))
			}
			err := candidate.reconnectCB()
			if err == nil {
				rp.RemoveReport(candidate.address)
			} else {
				rp.ReportProvider(candidate.address, 0, 1, nil) // add a disconnection
			}
			utils.LavaFormatDebug("reconnect attempt", utils.Attribute{Key: "provider", Value: candidate.address}, utils.Attribute{Key: "success", Value: err == nil})
		}
	}
}

func (rp *ReportedProviders) AppendReport(report metrics.ReportsRequest) {
	if rp == nil || rp.reporter == nil {
		return
	}
	if debugReportedProviders {
		utils.LavaFormatDebug("[debugReportedProviders] Sending report on provider", utils.LogAttr("provider", report.Provider))
	}
	rp.reporter.AppendReport(report)
}

func NewReportedProviders(reporter metrics.Reporter) *ReportedProviders {
	rp := &ReportedProviders{addedToPurgeAndReport: map[string]*ReportedProviderEntry{}, reporter: reporter}
	go func() {
		ticker := time.NewTicker(ReconnectCandidateTime)
		defer ticker.Stop()
		for range ticker.C {
			rp.ReconnectProviders()
		}
	}()
	return rp
}
