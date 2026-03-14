package mutexes

import "sync"

// TODO: Implement Counter using mutexes
// Read README.md for the instructions

// Counter tracks produced and consumed items safely across goroutines
type Counter struct {
	// 1. Declare a mutex to protect concurrent access
	mu sync.Mutex

	// 2. items stores the current count (produced - consumed)
	items int
}

// Produce adds items to the counter
func (c *Counter) Produce(amount int) {
	// 1. Lock the mutex before modifying shared state
	c.mu.Lock()

	// 2. Ensure the mutex is unlocked when the function exits
	defer c.mu.Unlock()

	// 3. Add the produced amount to the counter
	c.items += amount
}

// Consume removes items from the counter
func (c *Counter) Consume(amount int) {
	// 1. Lock the mutex before modifying shared state
	c.mu.Lock()

	// 2. Ensure the mutex is unlocked when the function exits
	defer c.mu.Unlock()

	// 3. Subtract the consumed amount from the counter
	c.items -= amount
}

// GetCount returns the current item count
func (c *Counter) GetCount() int {
	// 1. Lock the mutex before reading shared state
	c.mu.Lock()

	// 2. Ensure the mutex is unlocked when the function exits
	defer c.mu.Unlock()

	// 3. Return the current value of items
	return c.items
}
