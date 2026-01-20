package workerpool

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the worker-pool Quest 🎉")
	}
	os.Exit(code)
}

// helper to generate test input
func generateInput(n int) []string {
	input := make([]string, n)
	for i := 0; i < n; i++ {
		input[i] = "job" + strconv.Itoa(i+1)
	}
	return input
}

func TestWorkerPool(t *testing.T) {
	input := generateInput(100)

	results := WorkerPool(input)

	if len(results) != len(input) {
		t.Fatalf("expected %d results, got %d", len(input), len(results))
	}
	cnt := make(map[int]int)

	for _, r := range results {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			t.Fatalf("result missing worker ID: %s", r)
		}
		job, workerID := parts[0], parts[1]

		if !strings.HasPrefix(job, "JOB") {
			t.Fatalf("job not uppercased: %s", job)
		}

		id, err := strconv.Atoi(workerID)
		if err != nil {
			t.Fatalf("worker ID is not a number: %s", workerID)
		}

		if id < 1 || id > WORKER_COUNT {
			t.Fatalf("worker ID out of range: %d", id)
		}
		cnt[id]++
	}
	for id, c := range cnt {
		if c == 0 {
			t.Fatalf("worker %d did not process any jobs", id)
		}
	}
}
