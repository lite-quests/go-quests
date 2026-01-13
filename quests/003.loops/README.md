# Go Loops

## Concept
Loops are the primary mechanism for iteration in Go. The `for` loop is the only looping construct, supporting standard C-style loops, while-style loops, and range-based iteration over collections.

## References
- https://gobyexample.com/for
- https://gobyexample.com/range
- https://go.dev/tour/flowcontrol/1

## Quest

### Objective
Implement two functions: `SumEvenNumbers` to aggregate even integers, and `KeepOnlyConsonants` to filter strings within a slice.

### Requirements

#### `SumEvenNumbers`
- Function: `SumEvenNumbers(n int) int`
- Iterate from `1` to `n` (inclusive).
- Sum all even numbers.
- Return `0` if `n <= 0`.

#### `KeepOnlyConsonants`
- Function: `KeepOnlyConsonants(strs []string) []string`
- Return a new slice containing processed strings.
- For each string:
    - Remove all vowels (`a, e, i, o, u`), case-insensitive.
    - Keep all other characters (consonants, numbers, symbols).
- **Filtering Rule**: If a processed string becomes empty (i.e., it contained only vowels or was already empty), **exclude it** from the result slice.

### Inputs
- `SumEvenNumbers`: `n` (int)
- `KeepOnlyConsonants`: `strs` ([]string)

### Outputs
- `SumEvenNumbers`: `int` (sum of evens)
- `KeepOnlyConsonants`: `[]string` (filtered slice)

### Examples
- `SumEvenNumbers(10)` → `30` (2+4+6+8+10)
- `KeepOnlyConsonants(["hello", "aeiou", "world"])` → `["hll", "wrld"]`
- `KeepOnlyConsonants(["a", "b"])` → `["b"]`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/3.loops
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestSumEvenNumbers
--- PASS: TestSumEvenNumbers (0.00s)
=== RUN   TestKeepOnlyConsonants
--- PASS: TestKeepOnlyConsonants (0.00s)
PASS
```
