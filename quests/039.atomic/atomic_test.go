package atomic

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
		println(colorGreen, "Success! Completed the Atomic Operations Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestIncrement(t *testing.T) {
	c := &Counter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if got := c.Load(); got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
}

func TestDecrement(t *testing.T) {
	c := &Counter{}
	c.Increment()
	c.Increment()
	c.Decrement()
	if got := c.Load(); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}

func TestReset(t *testing.T) {
	c := &Counter{}
	c.Increment()
	c.Increment()
	c.Reset()
	if got := c.Load(); got != 0 {
		t.Errorf("expected 0 after reset, got %d", got)
	}
}

func TestCompareAndSwap_Success(t *testing.T) {
	c := &Counter{}
	c.Increment() // value = 1

	swapped := c.CompareAndSwap(1, 42)
	if !swapped {
		t.Error("expected swap to succeed")
	}
	if got := c.Load(); got != 42 {
		t.Errorf("expected 42 after CAS, got %d", got)
	}
}

func TestCompareAndSwap_Failure(t *testing.T) {
	c := &Counter{}
	c.Increment() // value = 1

	swapped := c.CompareAndSwap(99, 42) // wrong old value
	if swapped {
		t.Error("expected swap to fail")
	}
	if got := c.Load(); got != 1 {
		t.Errorf("expected value unchanged at 1, got %d", got)
	}
}

// TestConcurrentIncrement verifies the counter is safe under concurrent access.
func TestConcurrentIncrement(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	goroutines := 100
	increments := 1000

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				c.Increment()
			}
		}()
	}

	wg.Wait()

	want := int64(goroutines * increments)
	if got := c.Load(); got != want {
		t.Errorf("expected %d, got %d (race condition?)", want, got)
	}
}
