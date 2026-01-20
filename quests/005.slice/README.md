# Go Slices

## Concept
Slices are dynamic views into arrays. They are the most common data structure for collections in Go, providing powerful features for filtering, mapping, and resizing strings or lists of data.

## References
- https://gobyexample.com/slices
- https://gobyexample.com/range
- https://go.dev/blog/slices-intro

## Quest

### Objective
Implement `ProcessScores` to filter, normalize, and adjust a list of integer scores according to specific business rules.

### Requirements
- Function: `ProcessScores(scores []int) []int`
- **Output**: Return a **new slice**. Do not mutate the input slice.
- **Processing Chain (Order Matters)**:
    1. **Filter**: Remove invalid scores (`score < 0` or `score > 100`).
    2. **Normalize**: Set any failing score (`< 40`) to `40`.
    3. **Bonus**: If the count of resulting *valid* scores is **greater than 5**, add `5` to every score.
    4. **Cap**: Ensure no score exceeds `100` after the bonus application.
- Preservation: Maintain the relative order of scores.

### Inputs
- `scores`: `[]int` (list of raw scores)

### Outputs
- `[]int`: Processed list of scores.

### Examples
- `[30, 50, 110, -5, 80]`
  - Filter: `[30, 50, 80]` (-5, 110 removed)
  - Normalize: `[40, 50, 80]`
  - Count (3) <= 5: No bonus.
  - Result: `[40, 50, 80]`

- `[10, 20, 30, 40, 50, 60, 70]`
  - Filter: All keep.
  - Normalize: `[40, 40, 40, 40, 50, 60, 70]`
  - Count (7) > 5: Add 5 -> `[45, 45, 45, 45, 55, 65, 75]`
  - Cap: None exceed 100.
  - Result: `[45, 45, 45, 45, 55, 65, 75]`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/005.slice
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestProcessScores
--- PASS: TestProcessScores (0.00s)
PASS
```
