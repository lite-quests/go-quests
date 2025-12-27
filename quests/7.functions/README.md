# Go Functions Quest — Multiple Returns & Variadic

Your task is to implement **functions that return multiple values** and **functions that take a variable number of arguments**.

This quest focuses on:

- Returning multiple values including errors
- Using variadic functions
- Iterating over slices
- Returning derived results based on input
- Handling invalid input properly

---

## Reference

- [Multiple Return Values](https://gobyexample.com/multiple-return-values)
- [Variadic Functions](https://gobyexample.com/variadic-functions)

---

## Functions to Implement

```go
func Divide(a, b int) (int, error)
func SumAll(nums ...int) int
func MaxMin(nums ...int) (int, int, error)
func ConcatAll(sep string, strs ...string) string
```

---

## Description

You are building **utility functions** that demonstrate common Go function patterns:

1. Returning multiple values
2. Returning `error` when inputs are invalid
3. Accepting a variable number of arguments
4. Returning computed results from inputs

---

## Important Note on Errors

Whenever a function signature includes `error`, you **must return a proper `error` type**, not a string or `nil` incorrectly.

Use:

```go
import "errors"

errors.New("your message")   // simple static error
```

or

```go
import "fmt"

fmt.Errorf("formatted message %d", value)  // dynamic error
```

**Do NOT** return a string as an error:

```go
return 0, "cannot divide by zero" // ❌ INVALID
```

**Do NOT** return `nil` when an error is expected for invalid input.

---

## Function Behavior (Explicit)

### 1. `Divide`

```go
func Divide(a, b int) (int, error)
```

- Return `a / b` if `b != 0`
- Return a proper `error` if `b == 0`

Example:

```go
Divide(10, 2) // 5, nil
Divide(10, 0) // 0, error
```

---

### 2. `SumAll`

```go
func SumAll(nums ...int) int
```

- Accept any number of integers
- Return their sum
- Return `0` if no arguments are given

Example:

```go
SumAll(1, 2, 3, 4) // 10
SumAll()           // 0
```

---

### 3. `MaxMin`

```go
func MaxMin(nums ...int) (int, int, error)
```

- Accept any number of integers
- Return `(max, min, nil)` if at least one number is provided
- Return `(0, 0, error)` if no numbers are provided

Important:

- Use `range` to iterate over `nums`
- Update `max` and `min` during iteration

Example:

```go
MaxMin(1, 5, 3)  // 5, 1, nil
MaxMin()         // 0, 0, error
```

---

### 4. `ConcatAll`

```go
func ConcatAll(sep string, strs ...string) string
```

- Accept a **separator** and variable number of strings
- Return a single string where all inputs are joined using `sep`
- Return empty string if no strings are provided

Example:

```go
ConcatAll("-", "a", "b", "c") // "a-b-c"
ConcatAll(",", )               // ""
```

---

## Requirements

1. Use multiple return values for results and errors
2. Use variadic arguments for `SumAll`, `MaxMin`, and `ConcatAll`
3. Use `range` to iterate over slices
4. Properly handle empty or invalid inputs by returning errors
5. Do not panic for invalid input
6. Always return computed results, not internal state

---

## Common Beginner Mistakes (Tests Will Catch)

- Returning a string instead of an `error`
- Forgetting to return `error` on invalid input
- Using fixed arguments instead of variadic
- Assuming slice has elements without checking
- Not using `range` when iterating
- Returning zero values incorrectly for empty input

---

## Key Ideas (Beginner-Friendly)

- Go supports multiple return values naturally
- Errors are just another return value
- Variadic functions accept `...type` arguments
- Iteration using `range` works on variadic slices
- Always handle empty or invalid inputs carefully
- Use `errors.New()` or `fmt.Errorf()` to create errors

---

## Run Tests

```bash
go test ./quests/7.functions -v
```
