# Sorting

## Concept

Sorting is one of the most common operations in programming. Go's `sort` package provides built-in functions for sorting slices of common types, and a flexible `sort.Slice` for custom sorting logic.

Key functions:

- **`sort.Ints(s []int)`** — sorts a slice of ints in ascending order, in-place.
- **`sort.Strings(s []string)`** — sorts a slice of strings lexicographically, in-place.
- **`sort.Slice(s any, less func(i, j int) bool)`** — sorts any slice using a custom comparator. The `less` function receives two indices and returns `true` if element `i` should come before element `j`.
- **`sort.IntsAreSorted(s []int) bool`** — returns `true` if the slice is already sorted in ascending order.

```go
nums := []int{3, 1, 2}
sort.Ints(nums)
// nums: [1, 2, 3]

people := []Person{{Name: "Bob", Age: 25}, {Name: "Alice", Age: 30}}
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
// people: [{Bob 25} {Alice 30}]
```

## References

- [Go by Example: Sorting](https://gobyexample.com/sorting)
- [Go by Example: Sorting by Functions](https://gobyexample.com/sorting-by-functions)
- [pkg.go.dev/sort](https://pkg.go.dev/sort)

## Quest

### Objective

Implement five sorting functions covering built-in sort helpers and custom comparators.

### Requirements

#### 1. `SortInts(nums []int)`
- Sort the slice of ints in ascending order, in-place.
- Use `sort.Ints`.

#### 2. `SortStrings(strs []string)`
- Sort the slice of strings in ascending (lexicographic) order, in-place.
- Use `sort.Strings`.

#### 3. `SortByLength(strs []string)`
- Sort the slice of strings by length (shortest first), in-place.
- Use `sort.Slice` with a custom comparator.

#### 4. `SortByAge(people []Person)`
- A `Person` struct with `Name string` and `Age int` is already defined.
- Sort the slice by `Age` in ascending order, in-place.
- Use `sort.Slice` with a custom comparator.

#### 5. `IsSorted(nums []int) bool`
- Return `true` if the int slice is already sorted in ascending order.
- Use `sort.IntsAreSorted`.

### Examples

```go
nums := []int{5, 2, 8, 1}
SortInts(nums)
// nums: [1, 2, 5, 8]

strs := []string{"banana", "fig", "kiwi"}
SortByLength(strs)
// strs: ["fig", "kiwi", "banana"]

people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
SortByAge(people)
// people: [{Bob 25} {Alice 30}]

IsSorted([]int{1, 2, 3}) // true
IsSorted([]int{3, 1, 2}) // false
```

## Tips

- `sort.Slice` sorts in-place — no need to return anything.
- The `less(i, j int) bool` function should return `true` when element at index `i` should come **before** element at index `j`.
- `sort.Ints` and `sort.Strings` are just convenience wrappers — under the hood they use `sort.Slice` too.

## Testing

```bash
go test -v ./quests/036.sorting
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestSortInts
--- PASS: TestSortInts (0.00s)
=== RUN   TestSortInts_AlreadySorted
--- PASS: TestSortInts_AlreadySorted (0.00s)
=== RUN   TestSortStrings
--- PASS: TestSortStrings (0.00s)
=== RUN   TestSortByLength
--- PASS: TestSortByLength (0.00s)
=== RUN   TestSortByAge
--- PASS: TestSortByAge (0.00s)
=== RUN   TestIsSorted_True
--- PASS: TestIsSorted_True (0.00s)
=== RUN   TestIsSorted_False
--- PASS: TestIsSorted_False (0.00s)
=== RUN   TestIsSorted_SingleElement
--- PASS: TestIsSorted_SingleElement (0.00s)
PASS
```
