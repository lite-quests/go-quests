package atomic

import "sync/atomic"

// TODO: Implement the functions below.
// Read README.md for the instructions.

// Counter is a thread-safe counter backed by an atomic int64.
type Counter struct {
	value int64
}

// Increment atomically adds 1 to the counter.
func (c *Counter) Increment() {
	// TODO: Implement using atomic.AddInt64
	_ = atomic.AddInt64 // hint
}

// Decrement atomically subtracts 1 from the counter.
func (c *Counter) Decrement() {
	// TODO: Implement using atomic.AddInt64
}

// Load atomically reads and returns the current counter value.
func (c *Counter) Load() int64 {
	// TODO: Implement using atomic.LoadInt64
	return 0
}

// Reset atomically sets the counter value to 0.
func (c *Counter) Reset() {
	// TODO: Implement using atomic.StoreInt64
}

// CompareAndSwap atomically sets the counter to newVal only if its current
// value equals oldVal. Returns true if the swap was performed.
func (c *Counter) CompareAndSwap(oldVal, newVal int64) bool {
	// TODO: Implement using atomic.CompareAndSwapInt64
	return false
}
