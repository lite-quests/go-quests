# Go Functions

## Concept
Functions in Go are first-class citizens. Go supports multiple return values, named return parameters, and variadic functions (functions that accept a variable number of arguments). Error handling is typically done by returning an `error` as the last return value.

## References
- https://gobyexample.com/multiple-return-values
- https://gobyexample.com/variadic-functions
- https://go.dev/tour/basics/6

## Quest

### Objective
Implement four utility functions demonstrating multiple return values, variadic parameters, and error handling.

### Requirements

#### `Divide`
- Function: `Divide(a, b int) (int, error)`
- Return `a / b` and `nil` error if `b` is not zero.
- Return `0` and a non-nil error if `b` is zero.

#### `SumAll`
- Function: `SumAll(nums ...int) int`
- Return the sum of all integers passed as arguments.
- Return `0` if no arguments are provided.

#### `MaxMin`
- Function: `MaxMin(nums ...int) (int, int, error)`
- Return the maximum, minimum, and `nil` error if arguments are provided.
- Return `0, 0` and a non-nil error if no arguments are provided.

#### `ConcatAll`
- Function: `ConcatAll(sep string, strs ...string) string`
- Return a single string joining all `strs` with `sep`.
- Return an empty string if no strings are provided.

### Inputs
- `Divide`: Two integers.
- `SumAll`: Variadic integers.
- `MaxMin`: Variadic integers.
- `ConcatAll`: Separator string and variadic strings.

### Outputs
- See requirements for return types (values and errors).

### Examples
- `Divide(10, 2)` → `5, nil`
- `Divide(10, 0)` → `0, error`
- `SumAll(1, 2, 3)` → `6`
- `MaxMin(1, 5, 3)` → `5, 1, nil`
- `ConcatAll("-", "a", "b")` → `"a-b"`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/7.functions
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestFunctions
--- PASS: TestFunctions (0.00s)
PASS
```
