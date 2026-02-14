# Mutexes

## Concept

Mutexes (mutual exclusion locks) provide safe access to shared data across multiple goroutines. When multiple goroutines need to read and write the same data, a mutex ensures only one goroutine can access it at a time, preventing race conditions. Go's `sync.Mutex` has two main methods: `Lock()` to acquire exclusive access and `Unlock()` to release it.

## References

- https://gobyexample.com/mutexes
- https://gobyexample.com/atomic-counters
- https://go.dev/tour/concurrency/9

## Quest

### Objective

Implement a thread-safe `Counter` that tracks items produced and consumed. Multiple goroutines will increment and decrement the counter concurrently. Without proper synchronization, concurrent access causes race conditions and incorrect counts. Your task is to use a mutex to make the counter safe for concurrent use.

### Requirements

#### `Counter` struct

```go
type Counter struct {
    mu    sync.Mutex
    items int
}
```

- `items`: Current count (produced - consumed)
- `mu`: Mutex to protect concurrent access

#### `Produce` method

```go
func (c *Counter) Produce(amount int)
```

- Add `amount` to the counter
- Must be thread-safe using the mutex
- Lock before modifying, unlock after

#### `Consume` method

```go
func (c *Counter) Consume(amount int)
```

- Subtract `amount` from the counter
- Must be thread-safe using the mutex
- Lock before modifying, unlock after

#### `GetCount` method

```go
func (c *Counter) GetCount() int
```

- Return current count
- Must be thread-safe using the mutex
- Lock before reading, unlock after

### Inputs

- Methods accept `amount int` for produce/consume operations

### Outputs

- `GetCount()` returns the current count as `int`

**Why Mutexes Matter:**

Without the mutex, concurrent goroutines would race to read and write `items`, causing:

- Lost updates (two goroutines read same value, both increment, one write overwrites the other)
- Incorrect final count
- Unpredictable behavior

With the mutex, each operation completes atomically before another can start.

## Testing

To run the tests:

```bash
go test -v ./quests/021.mutexes
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestCounter
=== RUN   TestCounter/concurrent_producers_and_consumers
--- PASS: TestCounter (0.00s)
    --- PASS: TestCounter/concurrent_producers_and_consumers (0.00s)
PAS
```

## Bonus Challenge

Try running the tests with the race detector to verify thread safety:

```bash
go test -race -v ./quests/021.mutexes
```

This detects race conditions even if tests pass. A properly implemented solution should show no race warnings.
