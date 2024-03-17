// ============= GET =============
// if seen block >= requested block && requested_block > 0 -> no need to fetch seen block in advance of get.
// if requested_block < 0 -> parse block, replace requested block,
//     * check if seen_block >= replaced_requested_block -> no need to fetch seen and get.

// else fetch seen

// if seen >= requested_block -> seen_for_hash = requested (for key calculation)
// if seen < requested_block, seen_for_hash = 0

// fetch with the new key, check if result.seen_block >= my_seen ==> return hit
// else, the result is bad. cache miss.

// ============ SET =========
// if seen >= requested_block -> seen_for_hash = requested (for key calculation)
// if seen < requested_block, seen_for_hash = 0

package cache

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	sdkerrors "cosmossdk.io/errors"
	"github.com/dgraph-io/ristretto"
	"github.com/lavanet/lava/protocol/lavaprotocol"
	"github.com/lavanet/lava/protocol/parser"
	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	NotFoundError     = sdkerrors.New("Cache miss", 1, "cache entry for specific block and request wasn't found")                                                   // client could'nt connect to any provider.
	HashMismatchError = sdkerrors.New("Cache hit but hash mismatch", 2, "cache entry for specific block and request had a mismatching hash stored")                 // client could'nt connect to any provider.
	EntryTypeError    = sdkerrors.New("Cache hit but entry is a different object", 3, "cache entry for specific block and request had a mismatching object stored") // client could'nt connect to any provider.
)

const (
	DbValueConfirmationAttempts = 5
	SEP                         = ";"
)

type RelayerCacheServer struct {
	pairingtypes.UnimplementedRelayerCacheServer
	CacheServer *CacheServer
	cacheHits   uint64
	cacheMisses uint64
}

type CacheValue struct {
	Response         pairingtypes.RelayReply
	Hash             []byte
	OptionalMetadata []pairingtypes.Metadata
}

func (cv *CacheValue) ToCacheReply() *pairingtypes.CacheRelayReply {
	return &pairingtypes.CacheRelayReply{
		Reply:            &cv.Response,
		OptionalMetadata: cv.OptionalMetadata,
	}
}

func (cv *CacheValue) Cost() int64 {
	return int64(len(cv.Response.Data))
}

type LastestCacheStore struct {
	latestBlock          int64
	latestExpirationTime time.Time
}

func (cv *LastestCacheStore) Cost() int64 {
	return 8 + 16
}

func (s *RelayerCacheServer) getSeenBlockForSharedStateMode(chainId string, sharedStateId string) int64 {
	if sharedStateId != "" {
		id := latestBlockKey(chainId, sharedStateId)
		value, found := getNonExpiredFromCache(s.CacheServer.finalizedCache, id)
		if !found {
			utils.LavaFormatInfo("Failed fetching state from cache for this user id", utils.LogAttr("id", id))
			return 0 // we cant set the seen block in this case it will be returned 0 and wont be used in the consumer side.
		}
		utils.LavaFormatInfo("getting seen block cache", utils.LogAttr("id", id), utils.LogAttr("value", value))
		if cacheValue, ok := value.(int64); ok {
			return cacheValue
		}
		utils.LavaFormatFatal("Failed converting cache result to int64", nil, utils.LogAttr("value", value))
	}
	return 0
}

func (s *RelayerCacheServer) GetRelay(ctx context.Context, relayCacheGet *pairingtypes.RelayCacheGet) (*pairingtypes.CacheRelayReply, error) {
	cacheReply := &pairingtypes.CacheRelayReply{}
	var cacheReplyTmp *pairingtypes.CacheRelayReply
	var err error
	var seenBlock int64

	originalRequestedBlock := relayCacheGet.RequestedBlock // save requested block prior to swap
	if originalRequestedBlock < 0 {                        // we need to fetch stored latest block information.
		getLatestBlock := s.getLatestBlock(latestBlockKey(relayCacheGet.ChainId, ""))
		relayCacheGet.RequestedBlock = lavaprotocol.ReplaceRequestedBlock(originalRequestedBlock, getLatestBlock)
	}

	utils.LavaFormatDebug("Got Cache Get", utils.Attribute{Key: "request_hash", Value: string(relayCacheGet.RequestHash)},
		utils.Attribute{Key: "finalized", Value: relayCacheGet.Finalized},
		utils.Attribute{Key: "requested_block", Value: originalRequestedBlock},
		utils.Attribute{Key: "block_hash", Value: relayCacheGet.BlockHash},
		utils.Attribute{Key: "requested_block_parsed", Value: relayCacheGet.RequestedBlock},
		utils.Attribute{Key: "seen_block", Value: relayCacheGet.SeenBlock},
	)
	if relayCacheGet.RequestedBlock >= 0 { // we can only fetch
		// check seen block is larger than our requested block, we don't need to fetch seen block prior as its already larger than requested block
		if relayCacheGet.SeenBlock >= relayCacheGet.RequestedBlock {
			waitGroup := sync.WaitGroup{}
			waitGroup.Add(2) // currently we have two groups getRelayInner and getSeenBlock
			// fetch all reads at the same time.
			go func() {
				defer waitGroup.Done()
				cacheReplyTmp, err = s.getRelayInner(relayCacheGet)
				if cacheReplyTmp != nil {
					cacheReply = cacheReplyTmp // set cache reply only if its not nil, as we need to store seen block in it.
				}
			}()
			go func() {
				defer waitGroup.Done()
				// set seen block if required
				seenBlock = s.getSeenBlockForSharedStateMode(relayCacheGet.ChainId, relayCacheGet.SharedStateId)
			}()
			// wait for all reads to complete before moving forward
			waitGroup.Wait()
		} else {
			// our seen block might change our cache key value when shared state is enabled, we need to fetch it prior to moving forward
			// fetch seen block prior to cache
			seenBlock = s.getSeenBlockForSharedStateMode(relayCacheGet.ChainId, relayCacheGet.SharedStateId)
			if seenBlock > relayCacheGet.SeenBlock {
				relayCacheGet.SeenBlock = seenBlock // update state.
			}
			// now that we have our updated seen block state we can
			cacheReplyTmp, err = s.getRelayInner(relayCacheGet)
			if cacheReplyTmp != nil {
				cacheReply = cacheReplyTmp // set cache reply only if its not nil, as we need to store seen block in it.
			}
		}

		// set seen block.
		if seenBlock > cacheReply.SeenBlock {
			cacheReply.SeenBlock = seenBlock
		}
	} else {
		// set the error so cache miss will trigger.
		err = utils.LavaFormatDebug("Requested block is invalid",
			utils.LogAttr("requested block", relayCacheGet.RequestedBlock),
			utils.LogAttr("request_hash", string(relayCacheGet.RequestHash)),
		)
	}

	// add prometheus metrics asynchronously
	go func() {
		cacheMetricsContext, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		var hit bool
		if err != nil {
			s.cacheMiss(cacheMetricsContext, err)
		} else {
			hit = true
			s.cacheHit(cacheMetricsContext)
		}
		s.CacheServer.CacheMetrics.AddApiSpecific(originalRequestedBlock, relayCacheGet.ChainId, hit)
	}()
	return cacheReply, err
}

// formatHashKey formats the hash key by adding latestBlock and seenBlock information.
// If seenBlock is greater than or equal to latestBlock, seenBlock is set to latestBlock for key calculation,
// otherwise, it's set to 0.
func (s *RelayerCacheServer) formatHashKey(hash []byte, parsedRequestedBlock int64, seenBlock int64) []byte {
	// Handle seenBlock according to the specified rules
	if seenBlock >= parsedRequestedBlock {
		seenBlock = parsedRequestedBlock
	} else {
		seenBlock = 0
	}
	// Append the latestBlock and seenBlock directly to the hash using little-endian encoding
	hash = binary.LittleEndian.AppendUint64(hash, uint64(parsedRequestedBlock))
	hash = binary.LittleEndian.AppendUint64(hash, uint64(seenBlock))
	return hash
}

func (s *RelayerCacheServer) getRelayInner(relayCacheGet *pairingtypes.RelayCacheGet) (*pairingtypes.CacheRelayReply, error) {
	// cache key is compressed from:
	// 1. Request hash including all the information inside RelayPrivateData (Salt can cause issues if not dealt with on consumer side.)
	// 2. chain-id (same requests for different chains should get unique results)
	// 3. seen block to distinguish between seen entries and unseen entries.
	cacheKey := s.formatHashKey(relayCacheGet.RequestHash, relayCacheGet.RequestedBlock, relayCacheGet.SeenBlock)
	cacheVal, cache_source, found := s.findInAllCaches(relayCacheGet.Finalized, cacheKey)
	// TODO: use the information when a new block is finalized
	if !found {
		return nil, NotFoundError
	}
	if cacheVal.Hash == nil {
		// if we didn't store a hash its also always a match
		utils.LavaFormatDebug("returning response", utils.Attribute{Key: "cache_source", Value: cache_source},
			utils.Attribute{Key: "hash", Value: "nil"},
			utils.Attribute{Key: "response_data", Value: parser.CapStringLen(string(cacheVal.Response.Data))},
		)
		return cacheVal.ToCacheReply(), nil
	}
	// entry found, now we check the hash requested and hash stored
	if bytes.Equal(cacheVal.Hash, relayCacheGet.BlockHash) {
		utils.LavaFormatDebug("returning response", utils.Attribute{Key: "cache_source", Value: cache_source},
			utils.Attribute{Key: "hash", Value: "match"},
			utils.Attribute{Key: "response_data", Value: parser.CapStringLen(string(cacheVal.Response.Data))},
		)
		return cacheVal.ToCacheReply(), nil
	}
	// TODO: handle case where we have hash stored and it became finalized
	return nil, HashMismatchError
}

func (s *RelayerCacheServer) performInt64WriteWithValidationAndRetry(
	getBlockCallback func() int64,
	setBlockCallback func(),
	newInfo int64,
) {
	existingInfo := getBlockCallback()
	// validate we have a newer block than the existing stored in the db.
	if existingInfo <= newInfo { // refreshes state even if its equal
		// for seen block we expire the entry after one hour otherwise this user will stay in the db for ever
		setBlockCallback()
		// a validation routine to make sure we don't have a race for the block rewrites as there are concurrent writes.
		// this will be solved once we implement a db with a queue but for now its a good enough solution.
		go func() {
			for i := 0; i < DbValueConfirmationAttempts; i++ {
				time.Sleep(time.Millisecond) // add some delay between read attempts
				currentInfo := getBlockCallback()
				if currentInfo > newInfo {
					return // there is a newer block stored we are no longer relevant we can just stop validating.
				}
				if currentInfo < newInfo {
					// other cache set raced us and we need to rewrite our value again as its a newer value
					setBlockCallback()
				}
			}
		}()
	}
}

// this method tries to set the seen block a few times until it succeeds. it will try to make sure its value is set
// to prevent race conditions when we have a few writes in the same time with older values we will try to set our value eventually
// if its the newest seen block
func (s *RelayerCacheServer) setSeenBlockOnSharedStateMode(chainId, sharedStateId string, seenBlock int64) {
	if sharedStateId == "" {
		return
	}
	key := latestBlockKey(chainId, sharedStateId)
	set := func() {
		s.CacheServer.finalizedCache.SetWithTTL(key, seenBlock, 0, s.CacheServer.ExpirationFinalized)
	}
	get := func() int64 {
		return s.getSeenBlockForSharedStateMode(chainId, sharedStateId)
	}
	s.performInt64WriteWithValidationAndRetry(get, set, seenBlock)
}

func (s *RelayerCacheServer) SetRelay(ctx context.Context, relayCacheSet *pairingtypes.RelayCacheSet) (*emptypb.Empty, error) {
	if relayCacheSet.RequestedBlock < 0 {
		return nil, utils.LavaFormatError("invalid relay cache set data, request block is negative", nil, utils.Attribute{Key: "requestBlock", Value: relayCacheSet.RequestedBlock})
	}
	cacheKey := s.formatHashKey(relayCacheSet.RequestHash, relayCacheSet.RequestedBlock, relayCacheSet.SeenBlock)
	cacheValue := formatCacheValue(relayCacheSet.Response, relayCacheSet.BlockHash, relayCacheSet.Finalized, relayCacheSet.OptionalMetadata)
	utils.LavaFormatDebug("Got Cache Set", utils.Attribute{Key: "cacheKey", Value: string(cacheKey)},
		utils.Attribute{Key: "finalized", Value: fmt.Sprintf("%t", relayCacheSet.Finalized)},
		utils.Attribute{Key: "requested_block", Value: relayCacheSet.RequestedBlock},
		utils.Attribute{Key: "response_data", Value: parser.CapStringLen(string(relayCacheSet.Response.Data))},
		utils.Attribute{Key: "requestHash", Value: string(relayCacheSet.BlockHash)})
	// finalized entries can stay there
	if relayCacheSet.Finalized {
		cache := s.CacheServer.finalizedCache
		cache.SetWithTTL(cacheKey, cacheValue, cacheValue.Cost(), s.CacheServer.ExpirationFinalized)
	} else {
		cache := s.CacheServer.tempCache
		cache.SetWithTTL(cacheKey, cacheValue, cacheValue.Cost(), s.getExpirationForChain(time.Duration(relayCacheSet.AverageBlockTime), relayCacheSet.BlockHash))
	}
	// Setting the seen block for shared state.
	// Getting the max block number between the seen block on the consumer side vs the latest block on the response of the provider
	latestKnownBlock := int64(math.Max(float64(relayCacheSet.Response.LatestBlock), float64(relayCacheSet.SeenBlock)))
	s.setSeenBlockOnSharedStateMode(relayCacheSet.ChainId, relayCacheSet.SharedStateId, latestKnownBlock)
	s.setLatestBlock(latestBlockKey(relayCacheSet.ChainId, ""), latestKnownBlock)
	return &emptypb.Empty{}, nil
}

func (s *RelayerCacheServer) Health(ctx context.Context, req *emptypb.Empty) (*pairingtypes.CacheUsage, error) {
	cacheHits := atomic.LoadUint64(&s.cacheHits)
	cacheMisses := atomic.LoadUint64(&s.cacheMisses)
	return &pairingtypes.CacheUsage{CacheHits: cacheHits, CacheMisses: cacheMisses}, nil
}

func (s *RelayerCacheServer) cacheHit(ctx context.Context) {
	atomic.AddUint64(&s.cacheHits, 1)
	s.PrintCacheStats(ctx, "[+] cache hit")
}

func (s *RelayerCacheServer) cacheMiss(ctx context.Context, errPrint error) {
	atomic.AddUint64(&s.cacheMisses, 1)
	s.PrintCacheStats(ctx, "[-] cache miss, error:"+errPrint.Error())
}

func (s *RelayerCacheServer) PrintCacheStats(ctx context.Context, desc string) {
	health, err := s.Health(ctx, nil)
	if err != nil {
		_ = utils.LavaFormatError("Failed to get health response", err)
	}
	_ = utils.LavaFormatDebug(desc,
		utils.Attribute{Key: "misses", Value: strconv.FormatUint(health.CacheMisses, 10)},
		utils.Attribute{Key: "hits", Value: strconv.FormatUint(health.CacheHits, 10)},
	)
}

func (s *RelayerCacheServer) getLatestBlockInner(key string) (latestBlock int64, expirationTime time.Time) {
	value, found := getNonExpiredFromCache(s.CacheServer.finalizedCache, key)
	if !found {
		return spectypes.NOT_APPLICABLE, time.Time{}
	}
	if cacheValue, ok := value.(LastestCacheStore); ok {
		return cacheValue.latestBlock, cacheValue.latestExpirationTime
	}
	utils.LavaFormatError("latestBlock value is not a LastestCacheStore", EntryTypeError, utils.Attribute{Key: "value", Value: fmt.Sprintf("%+v", value)})
	return spectypes.NOT_APPLICABLE, time.Time{}
}

func (s *RelayerCacheServer) getLatestBlock(key string) int64 {
	latestBlock, expirationTime := s.getLatestBlockInner(key)
	if latestBlock != spectypes.NOT_APPLICABLE && expirationTime.After(time.Now()) {
		return latestBlock
	}
	return spectypes.NOT_APPLICABLE
}

func (s *RelayerCacheServer) setLatestBlock(key string, latestBlock int64) {
	cacheStore := LastestCacheStore{latestBlock: latestBlock, latestExpirationTime: time.Now().Add(DefaultExpirationForNonFinalized)}
	utils.LavaFormatDebug("setting latest block", utils.Attribute{Key: "key", Value: key}, utils.Attribute{Key: "latestBlock", Value: latestBlock})
	set := func() {
		s.CacheServer.finalizedCache.Set(key, cacheStore, cacheStore.Cost()) // no expiration time
	}
	get := func() int64 {
		existingLatest, _ := s.getLatestBlockInner(key) // we need to bypass the expirationTimeCheck
		return existingLatest
	}
	s.performInt64WriteWithValidationAndRetry(get, set, latestBlock)
}

func (s *RelayerCacheServer) getExpirationForChain(averageBlockTimeForChain time.Duration, blockHash []byte) time.Duration {
	if blockHash != nil {
		// this means that this entry has a block hash, so we don't have to delete it quickly
		return s.CacheServer.ExpirationFinalized
	}
	// if there is no block hash, for non finalized we cant know if there was a fork, so we have to delete it as soon as we have new data
	// with the assumption new data should arrive by the arrival of a new block (average block time)
	return s.CacheServer.ExpirationForChain(averageBlockTimeForChain)
}

func getNonExpiredFromCache(c *ristretto.Cache, key interface{}) (value interface{}, found bool) {
	value, found = c.Get(key)
	if found {
		return value, true
	}
	return nil, false
}

func (s *RelayerCacheServer) findInAllCaches(finalized bool, cacheKey []byte) (retVal CacheValue, cacheSource string, found bool) {
	inner := func(finalized bool, cacheKey []byte) (interface{}, string, bool) {
		if finalized {
			cache := s.CacheServer.finalizedCache
			value, found := getNonExpiredFromCache(cache, cacheKey)
			if found {
				return value, "finalized_cache", true
			}
			// if a key is finalized still doesn't mean it wasn't set when unfinalized
			cache = s.CacheServer.tempCache
			value, found = getNonExpiredFromCache(cache, cacheKey)
			if found {
				return value, "temp_cache", true
			}
		} else {
			// if something isn't finalized now it was never finalized, but sometimes when we don't have information we try to get a non finalized entry when in fact its finalized
			cache := s.CacheServer.tempCache
			value, found := getNonExpiredFromCache(cache, cacheKey)
			if found {
				return value, "temp_cache", true
			}
			cache = s.CacheServer.finalizedCache
			value, found = getNonExpiredFromCache(cache, cacheKey)
			if found {
				return value, "finalized_cache", true
			}
		}

		return nil, "", false
	}

	value, cacheSource, found := inner(finalized, cacheKey)
	if !found {
		return CacheValue{}, "", false
	}
	if cacheVal, ok := value.(CacheValue); ok {
		return cacheVal, cacheSource, true
	}
	utils.LavaFormatError("entry in cache was not a CacheValue", EntryTypeError, utils.Attribute{Key: "entry", Value: fmt.Sprintf("%+v", value)})
	return CacheValue{}, "", false
}

func formatCacheValue(response *pairingtypes.RelayReply, hash []byte, finalized bool, optionalMetadata []pairingtypes.Metadata) CacheValue {
	response.Sig = []byte{} // make sure we return a signed value, as the output was modified by our outputParser
	if !finalized {
		// hash value is only used on non finalized entries to check for forks
		return CacheValue{
			Response:         *response,
			Hash:             hash,
			OptionalMetadata: optionalMetadata,
		}
	}
	// no need to store the hash value for finalized entries
	return CacheValue{
		Response:         *response,
		Hash:             nil,
		OptionalMetadata: optionalMetadata,
	}
}

// used both by shared-state id and provider address id. so we just call the 2nd arg unique id
// as it can be both provider address or the unique user id (ip + dapp id)
func latestBlockKey(chainID string, uniqueId string) string {
	// because we want to support coherence in providers
	return chainID + "_" + uniqueId
}
