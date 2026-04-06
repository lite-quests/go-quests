package atomic

import "sync/atomic"

// Counter is a thread-safe counter backed by an atomic int64.
type Counter struct {
	value int64
}

// Increment atomically adds 1 to the counter.
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Decrement atomically subtracts 1 from the counter.
func (c *Counter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

// Load atomically reads and returns the current counter value.
func (c *Counter) Load() int64 {
	return atomic.LoadInt64(&c.value)
}

// Reset atomically sets the counter value to 0.
func (c *Counter) Reset() {
	atomic.StoreInt64(&c.value, 0)
}

// CompareAndSwap atomically sets the counter to newVal only if current value equals oldVal.
func (c *Counter) CompareAndSwap(oldVal, newVal int64) bool {
	return atomic.CompareAndSwapInt64(&c.value, oldVal, newVal)
}
