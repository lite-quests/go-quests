# Go Generics

## Concept
Generics allow defining data structures and functions that work with any type, providing type safety without code duplication. Type parameters (e.g., `[T any]`) are specified at definition and instantiated at use.

## References
- https://gobyexample.com/generics
- https://go.dev/doc/tutorial/generics
- https://go.dev/blog/intro-generics

## Quest

### Objective
Implement a generic singly-linked list `List[T]` capable of storing elements of any type, with methods to add, remove, and retrieve elements.

### Requirements

#### Structs
- `List[T any]`: Contains `head` and `tail` pointers to `element[T]`.
- `element[T any]`: Contains `val` (type `T`) and `next` pointer.

#### Methods
1.  **`Push(v T)`**
    -   Add `v` to the **end** of the list.
    -   Update `head` and `tail` appropriately.
    -   Handle empty list case.

2.  **`Pop() (T, bool)`**
    -   Remove the **last** element (tail) from the list.
    -   Return the value and `true`.
    -   If list is empty, return zero value of `T` and `false`.
    -   Update `head` and `tail` (O(N) operation allows logical correctness for singly-linked).

3.  **`AllElements() []T`**
    -   Return a slice containing all elements in order from head to tail.
    -   Return an empty slice (not `nil`) if list is empty.

### Inputs
- `v`: Value of type `T` to push.

### Outputs
- `Pop`: Value of type `T` and boolean success flag.
- `AllElements`: Slice `[]T`.

### Examples
```go
lst := List[int]{}
lst.Push(10)
lst.Push(20)
// List: 10 -> 20

val, ok := lst.Pop()
// val: 20, ok: true
// List: 10

elems := lst.AllElements()
// elems: [10]
```

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/13.generics
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestElements_Int
--- PASS: TestElements_Int (0.00s)
=== RUN   TestElements_String
--- PASS: TestElements_String (0.00s)
PASS
```
