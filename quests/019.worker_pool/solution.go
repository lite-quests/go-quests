package workerpool

import "time"

const WORKER_COUNT = 5

func worker(id int, inputs <-chan string, outputs chan<- string) {
	// TODO: Implement the worker function
	// - read string from inputs channel
	// - convert to uppercase
	// - append -ID
	// - send result to outputs channel
	time.Sleep(time.Second) // optional simulated work
}

func WorkerPool(strs []string) []string {
	// TODO:
	// 1. Create inputs and outputs channels with capacity = len(strs)
	// 2. Start WORKER_COUNT workers
	// 3. Send all strings into inputs channel
	// 4. Close inputs channel
	// 5. Collect all results from outputs channel
	// 6. Return slice of processed strings
	return []string{}
}
