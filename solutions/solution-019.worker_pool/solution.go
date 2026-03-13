package solution019workerpool

import (
	"strconv"
	"strings"
)

const WORKER_COUNT = 5

func worker(id int, inputs <-chan string, outputs chan<- string) {
	for str := range inputs {
		result := strings.ToUpper(str) + "-" + strconv.Itoa(id)
		outputs <- result
	}
}

func WorkerPool(strs []string) []string {
	inputs := make(chan string, len(strs))
	outputs := make(chan string, len(strs))

	for i := 1; i <= WORKER_COUNT; i++ {
		go worker(i, inputs, outputs)
	}

	for _, s := range strs {
		inputs <- s
	}

	close(inputs)

	result := make([]string, 0, len(strs))

	for range strs {
		result = append(result, <-outputs)
	}

	return result
}
