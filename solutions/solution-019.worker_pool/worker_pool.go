package worker_pool

import (
	"strconv"
	"strings"
)

// TODO: Implement WorkerPool
// Read README.md for the instructions

const WORKER_COUNT = 5

// worker pulls strings from inputs, transforms them, and sends results to outputs.
// It exits when the inputs channel is closed.
func worker(id int, inputs <-chan string, outputs chan<- string) {
	// Range over inputs until the channel is closed by the caller
	for str := range inputs {
		// Uppercase the string and tag it with this worker's ID
		result := strings.ToUpper(str) + "-" + strconv.Itoa(id)
		outputs <- result
	}
}

func WorkerPool(strs []string) []string {
	// 1. Create buffered channels sized to input count to avoid blocking
	inputs := make(chan string, len(strs))
	outputs := make(chan string, len(strs))

	// 2. Spawn WORKER_COUNT goroutines — each pulls work from inputs concurrently
	for i := 1; i <= WORKER_COUNT; i++ {
		go worker(i, inputs, outputs)
	}

	// 3. Feed all input strings into the inputs channel
	for _, s := range strs {
		inputs <- s
	}

	// 4. Close inputs to signal workers there's no more work coming
	close(inputs)

	// 5. Collect exactly len(strs) results from outputs
	result := make([]string, 0, len(strs))
	for range strs {
		result = append(result, <-outputs)
	}

	return result
}
