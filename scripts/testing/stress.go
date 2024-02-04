package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

const baseURL = "https://g.w.lavanet.xyz:443/gateway/" // Base URL as a constant

var (
	requestsPerSecond int
	duration          int = 1 // Set default duration to 1
	blockID           int
	projectHash       string
	chainID           string // Variable to store the chain ID
	times             []time.Duration
	errors            []error // To track errors for ASCII table
	mutex             sync.Mutex
	client            *http.Client
)

// Payloads map for different chainIDs
var payloads = map[string]string{
	"near": `{"jsonrpc":"2.0","method":"block","params":{"block_id":%d},"id":1}`,
	"cos3": `{"jsonrpc":"2.0","method":"block","params":{"block_id":%d},"id":1}`,
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	client = &http.Client{}

	rootCmd := &cobra.Command{
		Use:   "stress",
		Short: "A stress test tool",
		Long: `A stress test tool to test the performance of the gateway by sending a specified number of requests per second (RPS) for a given duration.

Examples:
  # Stress test the "near" chain with 10 RPS for 5 seconds
  ./stress --block-id 1234 --project-hash b37d228eceac369d476ff0597e5bfc1f --rps 10 --duration 5

  # Stress test the "axelar" chain with 5 RPS for 3 seconds
  ./stress --block-id 5678 --project-hash b37d228eceac369d476ff0597e5bfc1f --rps 5 --duration 3`,
		Run: runStressTest,
	}

	rootCmd.Flags().IntVarP(&blockID, "block-id", "b", 0, "Block ID to request (mandatory)")
	rootCmd.Flags().IntVarP(&duration, "duration", "d", 1, "Duration in seconds for the test")
	rootCmd.Flags().IntVarP(&requestsPerSecond, "rps", "r", 1, "Number of requests per second (mandatory)")
	rootCmd.Flags().StringVarP(&projectHash, "project-hash", "p", "", "Project hash (mandatory)")
	rootCmd.Flags().StringVarP(&chainID, "chain-id", "c", "", "Chain ID (options: near, axelar)")

	rootCmd.MarkFlagRequired("block-id")
	rootCmd.MarkFlagRequired("rps")
	rootCmd.MarkFlagRequired("chain-id")
	rootCmd.MarkFlagRequired("project-hash")

	_, ok := payloads[chainID]
	if !ok && chainID != "" {
		fmt.Printf("Unsupported chain ID: %s\n", chainID)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func runStressTest(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	statusCodeDistribution := make(map[int]int)
	var totalRequests, successCount int
	currentBlockID := blockID // Start from the initial blockID

	for i := 0; i < duration; i++ {
		for j := 0; j < requestsPerSecond; j++ {
			wg.Add(1)
			go func(blockID int) {
				defer wg.Done()
				startTime := time.Now()
				statusCode, err := makeRequest(blockID)
				duration := time.Since(startTime)

				mutex.Lock()
				if err != nil {
					errors = append(errors, err)
				} else {
					if statusCode == http.StatusOK {
						successCount++
						times = append(times, duration)
					}
					statusCodeDistribution[statusCode]++
				}
				mutex.Unlock()
			}(currentBlockID)
			totalRequests++
			currentBlockID++ // Increment blockID for the next request
		}
		wg.Wait()
	}

	printSummary(duration, requestsPerSecond, totalRequests, successCount, statusCodeDistribution)
	printPercentiles()
	if len(errors) > 0 {
		printErrorSummary()
	}
}

func makeRequest(blockID int) (int, error) {
	fullURL := fmt.Sprintf("%s%s/rpc-http/%s", baseURL, chainID, projectHash)
	payload, ok := payloads[chainID]
	if !ok {
		return 0, fmt.Errorf("unsupported chain ID: %s", chainID)
	}
	formattedPayload := fmt.Sprintf(payload, blockID)

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(formattedPayload))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func printSummary(duration, rps, total, success int, distribution map[int]int) {
	fmt.Println("Test Summary")
	fmt.Println("------------")
	fmt.Printf("Duration: %ds | RPS: %d | Total Requests: %d | Successful Requests: %d\n", duration, rps, total, success)
	fmt.Println("\nHTTP Status Code Distribution")
	fmt.Println("----------------------------")
	fmt.Println("Status Code | Count")
	for code, count := range distribution {
		fmt.Printf("%11d | %5d\n", code, count)
	}
	fmt.Printf("Success Rate: %.2f%%\n", (float64(success)/float64(total))*100)
}

func printPercentiles() {
	fmt.Println("\nResponse Time Percentiles (ms)")
	fmt.Println("------------------------------")
	sortDurations(times)
	fmt.Println("P50 | P90 | P95")
	fmt.Printf("%3d | %3d | %3d\n", percentile(50), percentile(90), percentile(95))
}

func printErrorSummary() {
	fmt.Println("\nError Summary")
	fmt.Println("-------------")
	errorDistribution := make(map[string]int)
	for _, err := range errors {
		errorDistribution[err.Error()]++
	}

	for err, count := range errorDistribution {
		fmt.Printf("%-70s | %d\n", err, count) // Adjusted for alignment
	}
}

func sortDurations(durations []time.Duration) {
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})
}

func percentile(p int) int {
	index := (p * len(times) / 100) - 1
	if index < 0 {
		index = 0
	}
	return int(times[index].Milliseconds())
}
