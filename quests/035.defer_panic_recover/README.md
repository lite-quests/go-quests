# Defer, Panic & Recover

## Concept

Go has three built-in mechanisms for controlling unusual execution flow:

### defer
A `defer` statement schedules a function call to run just before the surrounding function returns — no matter how it returns (normally, via `return`, or via `panic`). Multiple defers in the same function run in **LIFO** (last-in, first-out) order, like a stack.

```go
func example() {
    defer fmt.Println("third") // runs last
    defer fmt.Println("second")
    defer fmt.Println("first") // runs first... wait, no — LIFO!
    // Output order: first → second → third? No!
    // Defers run in reverse: "first" deferred last runs first? No.
    // The LAST defer registered runs FIRST.
    // So output: first, second, third → reversed: third, second, first
}
```

The most common use of `defer` is resource cleanup — closing files, releasing locks, etc. — so you can write the cleanup right next to the open, and it always runs.

### panic
`panic` stops normal execution of the current goroutine. It unwinds the call stack, running any deferred functions along the way. If nothing recovers from it, the program crashes.

```go
panic("something went wrong") // stops execution
```

### recover
`recover` stops a panic and returns the value passed to `panic`. It only works inside a **deferred** function. If there's no panic, `recover` returns `nil`.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}()
panic("oops") // recovered above
```

## References

- [Go by Example: Defer](https://gobyexample.com/defer)
- [Go by Example: Panic](https://gobyexample.com/panic)
- [Go by Example: Recover](https://gobyexample.com/recover)
- [Effective Go: Defer, Panic, Recover](https://go.dev/doc/effective_go#defer)

## Quest

### Objective

Implement four small functions that demonstrate `defer` ordering, `panic`, `recover`, and deferred cleanup.

### Requirements

#### 1. `DeferOrder() []string`
- Create a `[]string` called `result`.
- Use `defer` to append `"first"`, `"second"`, and `"third"` to `result` (in that registration order).
- Return `result` after all defers have run.
- Expected return value: `["third", "second", "first"]` (LIFO order).

#### 2. `SafeDivide(a, b int) int`
- Return `a / b`.
- If `b == 0`, call `panic("division by zero")`.

#### 3. `RecoverDivide(a, b int) (result int, panicMsg string)`
- Call `SafeDivide(a, b)` inside a deferred recover.
- If a panic occurs, return `0` and the panic message as a string.
- If no panic, return the division result and an empty string.

#### 4. `Cleanup() []string`
- Create a `[]string` called `log`.
- Append `"open"` to `log`.
- Use `defer` to append `"close"` to `log`.
- Append `"work"` to `log`.
- Return `log` after all defers run.
- Expected return value: `["open", "work", "close"]`.

### Examples

```go
DeferOrder()
// returns: ["third", "second", "first"]

SafeDivide(10, 2)
// returns: 5

SafeDivide(10, 0)
// panics with: "division by zero"

RecoverDivide(10, 2)
// returns: 5, ""

RecoverDivide(10, 0)
// returns: 0, "division by zero"

Cleanup()
// returns: ["open", "work", "close"]
```

## Tips

- In `DeferOrder`, the defers run in reverse registration order — the last `defer` registered runs first.
- In `RecoverDivide`, the `defer` with `recover()` must be set up **before** calling `SafeDivide`.
- In `Cleanup`, `defer` guarantees `"close"` runs even if the function panics midway.
- `recover()` returns `any` — use a type assertion `r.(string)` to get the string message.

## Testing

```bash
go test -v ./quests/035.defer_panic_recover
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestDeferOrder
--- PASS: TestDeferOrder (0.00s)
=== RUN   TestSafeDivide_Normal
--- PASS: TestSafeDivide_Normal (0.00s)
=== RUN   TestSafeDivide_PanicsOnZero
--- PASS: TestSafeDivide_PanicsOnZero (0.00s)
=== RUN   TestRecoverDivide_Normal
--- PASS: TestRecoverDivide_Normal (0.00s)
=== RUN   TestRecoverDivide_PanicRecovered
--- PASS: TestRecoverDivide_PanicRecovered (0.00s)
=== RUN   TestCleanup
--- PASS: TestCleanup (0.00s)
PASS
```
