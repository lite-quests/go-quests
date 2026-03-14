package mutexes

import (
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the mutexes Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestCounter(t *testing.T) {
	t.Run("concurrent_producers_and_consumers", func(t *testing.T) {
		counter := &Counter{}
		var wg sync.WaitGroup

		// 10 producers, each producing 100 items
		producers := 10
		itemsPerProducer := 100

		// 5 consumers, each consuming 100 items
		consumers := 5
		itemsPerConsumer := 100

		expected := (producers * itemsPerProducer) - (consumers * itemsPerConsumer)

		// Start producers
		for i := 0; i < producers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < itemsPerProducer; j++ {
					counter.Produce(1)
				}
			}()
		}

		// Start consumers
		for i := 0; i < consumers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < itemsPerConsumer; j++ {
					counter.Consume(1)
				}
			}()
		}

		wg.Wait()

		if got := counter.GetCount(); got != expected {
			t.Errorf("expected %d, got %d (race condition detected - did you use mutex?)", expected, got)
		}
	})
}
