# Closures

## Concept

A **closure** is a function that "closes over" variables from its surrounding scope — it captures and remembers them even after the outer function has returned.

In Go, functions are first-class values. You can return a function from another function, and the returned function will carry its own private copy of any variables it captured.

```go
func makeCounter() func() int {
    count := 0           // this variable is captured
    return func() int {
        count++          // each call modifies the same `count`
        return count
    }
}

counter := makeCounter()
counter() // 1
counter() // 2
counter() // 3

// A second counter is completely independent
other := makeCounter()
other() // 1  (its own `count`, starting from 0)
```

Each call to `makeCounter()` creates a **new** closure with its **own** `count` variable. They don't share state.

This pattern is used everywhere: event handlers, middleware, factory functions, and stateful iterators.

## References

- [Go by Example: Closures](https://gobyexample.com/closures)
- [Go Tour: Function closures](https://go.dev/tour/moretypes/25)

## Quest

### Objective

Implement four factory functions that return closures, each demonstrating a different use of captured state.

### Requirements

#### 1. `MakeCounter() func() int`
- Returns a function with no arguments.
- Each call to the returned function increments an internal counter (starting at 0) and returns the new value.
- Multiple counters created by `MakeCounter()` must be independent.

#### 2. `MakeAdder(n int) func(int) int`
- Returns a function that adds `n` to its argument.
- `n` is captured from the outer function.

#### 3. `MakeMultiplier(factor int) func(int) int`
- Returns a function that multiplies its argument by `factor`.
- `factor` is captured from the outer function.

#### 4. `MakeAccumulator() func(int) int`
- Returns a function that maintains a running total.
- Each call adds the given value to the total and returns the new total.
- Multiple accumulators must be independent.

### Examples

```go
counter := MakeCounter()
counter() // 1
counter() // 2
counter() // 3

add5 := MakeAdder(5)
add5(3)  // 8
add5(10) // 15

double := MakeMultiplier(2)
double(4)  // 8
double(-3) // -6

acc := MakeAccumulator()
acc(5)  // 5
acc(3)  // 8
acc(-4) // 4
```

## Tips

- Declare the variable you want to capture **before** the `return func(...)` statement.
- The returned function modifies that variable directly — no need to pass it as an argument.
- Each call to the factory function (`MakeCounter`, `MakeAccumulator`) creates a brand new variable, so each returned closure is independent.

## Testing

```bash
go test -v ./quests/037.closures
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestMakeCounter
--- PASS: TestMakeCounter (0.00s)
=== RUN   TestMakeCounter_Independent
--- PASS: TestMakeCounter_Independent (0.00s)
=== RUN   TestMakeAdder
--- PASS: TestMakeAdder (0.00s)
=== RUN   TestMakeMultiplier
--- PASS: TestMakeMultiplier (0.00s)
=== RUN   TestMakeAccumulator
--- PASS: TestMakeAccumulator (0.00s)
=== RUN   TestMakeAccumulator_Independent
--- PASS: TestMakeAccumulator_Independent (0.00s)
PASS
```
