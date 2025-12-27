# Go Arrays & Slices Quest

Your task is to implement **slice-processing logic** using idiomatic Go.

This quest focuses on:

- Iterating over slices safely
- Building new slices (not mutating inputs)
- Index-based vs range-based loops
- Understanding slice length and order guarantees

---

## Reference

- [https://gobyexample.com/slices](https://gobyexample.com/slices)
- [https://gobyexample.com/arrays](https://gobyexample.com/arrays)
- [https://gobyexample.com/range](https://gobyexample.com/range)

---

## Function: `ProcessScores`

```go
func ProcessScores(scores []int) []int
```

---

## Description

You are given a slice of integers representing raw scores.

Your task is to return a **new slice** after applying a sequence of transformations.

The original slice **must not be modified**.

---

## Transformation Rules (Order Matters)

Apply the following steps **in order**:

---

### 1. Filter Invalid Scores

Remove any score that is:

- Less than `0`
- Greater than `100`

---

### 2. Normalize Failing Scores

For the remaining scores:

- Any score **below 40** must be replaced with `40`

---

### 3. Bonus Adjustment

- If the number of **valid scores** after filtering is **greater than 5**, add **+5** to every score
- Scores must be capped at `100` after adjustment

---

### 4. Preserve Order

The relative order of scores must remain unchanged throughout.

---

## Requirements

1. Use slices, not arrays
2. Do not mutate the input slice
3. Use at least one `append`
4. Use loops (`for` / `range`)

---

## Examples

### Example 1

```go
input  := []int{30, 50, 110, -5, 80}
output := []int{40, 50, 80}
```

Explanation:

- `110`, `-5` removed
- `30 → 40`
- Only 3 valid scores → no bonus

---

### Example 2

```go
input  := []int{10, 20, 30, 40, 50, 60, 70}
output := []int{45, 45, 45, 45, 55, 65, 75}
```

Explanation:

- All valid
- Scores below 40 normalized
- 7 valid scores → +5 bonus

---

### Example 3

```go
input  := []int{}
output := []int{}
```

---

## Common Mistakes (Tests Will Catch These)

- Mutating the input slice
- Applying bonus before normalization
- Forgetting to cap scores at 100
- Losing element order
- Using array instead of slice semantics

---

## Key Ideas

- Slices are views; copying matters
- Length determines logic, capacity determines performance
- Order-preserving transforms are common in real systems
- Build pipelines deliberately

---

## Run Tests

```bash
go test ./quests/5.slice -v
```
