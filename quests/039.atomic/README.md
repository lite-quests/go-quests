# Atomic Operations

## Concept

When multiple goroutines read and write the same variable concurrently, you get a **data race** — unpredictable, corrupted results. The usual fix is a `sync.Mutex`, but for simple numeric operations there's a lighter tool: `sync/atomic`.

Atomic operations are **indivisible** — no other goroutine can observe a half-completed read or write. They're faster than mutexes for simple counters and flags.

Key functions from `sync/atomic`:

| Function | What it does |
|---|---|
| `atomic.AddInt64(&v, delta)` | Atomically adds delta to v, returns new value |
| `atomic.LoadInt64(&v)` | Atomically reads v |
| `atomic.StoreInt64(&v, val)` | Atomically writes val to v |
| `atomic.CompareAndSwapInt64(&v, old, new)` | Sets v=new only if v==old, returns bool |

```go
var count int64

atomic.AddInt64(&count, 1)          // increment
atomic.AddInt64(&count, -1)         // decrement
val := atomic.LoadInt64(&count)     // read
atomic.StoreInt64(&count, 0)        // reset

// Compare-and-swap: only update if current value matches
swapped := atomic.CompareAndSwapInt64(&count, 5, 10)
```

## References

- [Go by Example: Atomic Counters](https://gobyexample.com/atomic-counters)
- [pkg.go.dev/sync/atomic](https://pkg.go.dev/sync/atomic)

## Quest

### Objective

Implement a thread-safe `Counter` type using `sync/atomic` operations.

### Requirements

A `Counter` struct with an `int64` field is already defined. Implement the following methods:

#### 1. `Increment()`
- Atomically add `1` to the counter.
- Use `atomic.AddInt64`.

#### 2. `Decrement()`
- Atomically subtract `1` from the counter.
- Use `atomic.AddInt64` with `-1`.

#### 3. `Load() int64`
- Atomically read and return the current value.
- Use `atomic.LoadInt64`.

#### 4. `Reset()`
- Atomically set the counter to `0`.
- Use `atomic.StoreInt64`.

#### 5. `CompareAndSwap(oldVal, newVal int64) bool`
- Atomically set the counter to `newVal` **only if** its current value equals `oldVal`.
- Return `true` if the swap happened, `false` otherwise.
- Use `atomic.CompareAndSwapInt64`.

### Why not just use `c.value++`?

Without atomics, concurrent increments cause a data race:

```
goroutine 1: reads value = 5
goroutine 2: reads value = 5   ← both read the same value!
goroutine 1: writes value = 6
goroutine 2: writes value = 6  ← one increment is lost
```

With `atomic.AddInt64`, the read-modify-write is a single indivisible operation.

### Examples

```go
c := &Counter{}
c.Increment()  // value: 1
c.Increment()  // value: 2
c.Decrement()  // value: 1
c.Load()       // 1
c.Reset()      // value: 0

c.Increment()                    // value: 1
c.CompareAndSwap(1, 42)          // true  → value: 42
c.CompareAndSwap(1, 99)          // false → value still 42
```

## Testing

```bash
go test -v ./quests/039.atomic
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestIncrement
--- PASS: TestIncrement (0.00s)
=== RUN   TestDecrement
--- PASS: TestDecrement (0.00s)
=== RUN   TestReset
--- PASS: TestReset (0.00s)
=== RUN   TestCompareAndSwap_Success
--- PASS: TestCompareAndSwap_Success (0.00s)
=== RUN   TestCompareAndSwap_Failure
--- PASS: TestCompareAndSwap_Failure (0.00s)
=== RUN   TestConcurrentIncrement
--- PASS: TestConcurrentIncrement (0.00s)
PASS
```
