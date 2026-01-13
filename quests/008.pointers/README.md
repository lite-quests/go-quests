# Go Pointers

## Concept
Pointers store the memory address of a value. They allow functions to mutate values stored in the caller's scope and avoid copying large data structures. Go passes arguments by value; to pass a reference, you must use a pointer.

## References
- https://gobyexample.com/pointers
- https://go.dev/tour/moretypes/1
- https://go.dev/doc/effective_go#pointers_vs_values

## Quest

### Objective
Modify the `main` function in `solution.go` to call the correct helper functions. You must achieve specific variable mutations by choosing between value-based and pointer-based functions.

### Requirements
- **Function**: `PointersQuest()`
- **Tasks**:
    1. **Task 1**: Change `x` (int) from `5` to `15`.
    2. **Task 2**: Change `y` (int) from `42` to `0`.
    3. **Task 3**: Swap `a` and `b` (ints) so `a` becomes `9` and `b` becomes `3`.
    4. **Task 4**: Append `100` to slice `s`.
- **Constraint**: You may **not** modify the helper functions (`addTenValue`, `addTenPointer`, etc.). You must call the existing functions with the correct arguments (values or pointers).

### Inputs
- None (The function logic is self-contained).

### Outputs
- The function prints the modified values to standard output:
    - `15`
    - `0`
    - `9 3`
    - `[1 2 3 100]`

### Examples
No direct function call examples; the quest is to implement the correct calls inside `PointersQuest`.

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/8.pointers
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestPointersQuest
--- PASS: TestPointersQuest (0.00s)
PASS
```
