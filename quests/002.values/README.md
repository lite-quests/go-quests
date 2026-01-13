# Go Values

## Concept
Go is a statically typed language with a rich set of value types. These include basic types (integers, floats, booleans, strings), composite types (arrays, slices, maps, structs), pointers, functions, and interfaces. Understanding how to instantiate and assign these values is fundamental.

## References
- https://go.dev/tour/basics/11
- https://gobyexample.com/values
- https://go.dev/ref/spec#Types

## Quest

### Objective
Implement the `BuildValues` function to construct and return a `Result` struct populated with specific values across various Go types.

### Requirements
- Function: `BuildValues() Result`
- Package: `values`
- Return a `Result` struct with the following fields explicitly set:
    - `Str`: `"go"`
    - `Int`: `42`
    - `Float`: `3.14`
    - `Bool`: `true`
    - `Array`: `[3]int` containing `1, 2, 3`
    - `Slice`: `[]int` containing `4, 5, 6, 7`
    - `Map`: `map[string]int` with keys `"apple"` mapping to `2` and `"banana"` mapping to `5`
    - `User`: `User` struct with `Name` `"Alice"` and `Age` `20`
    - `Ptr`: A pointer to an integer with value `10`
    - `AddFn`: A function that accepts two integers and returns their sum
    - `Any`: An `interface{}` holding the integer `100`
    - `ZeroMap`: An uninitialized map (nil)

### Inputs
None

### Outputs
- Type: `Result` (struct)
- Format: Struct containing specified fields.
- Expected behavior: The returned struct matches the requirements in the test suite.

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/2.values
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestBuildValues
--- PASS: TestBuildValues (0.00s)
PASS
```