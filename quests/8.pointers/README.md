# Go Pointers Quest — Values vs Memory Addresses

Your task is to understand and practice **how Go passes data to functions** and how **pointers allow functions to mutate caller-owned memory**.

This quest focuses on:

- The difference between passing **values** and **addresses**
- Using `&` to take an address
- Using `*` to read or modify a value at an address
- Predicting whether a function call mutates state
- Choosing the correct API based on intent

---

## What Is a Pointer (Conceptual)

In Go, every variable lives at a **memory address**.

```go
x := 10
```

Conceptually:

```
address ──► value
0x100    ──► 10
```

A **pointer** is a variable that stores that address.

```go
p := &x
```

Now:

- `x` is the value (`10`)
- `&x` is the address (`0x100`)
- `p` stores the address
- `*p` accesses the value at that address (`10`)

### Why Pointers Matter

By default, Go passes **copies** to functions.

```go
func inc(n int) {
	n++
}
```

This does not affect the caller.

Pointers allow a function to **operate on the original memory** instead of a copy.

```go
func inc(n *int) {
	*n++
}
```

This is the _only_ reason pointers exist in Go.

---

## Reference

- [Pointers](https://gobyexample.com/pointers)
- [Nil Pointers](https://gobyexample.com/nil)

---

## Provided Functions (Do NOT Modify)

```go
func addTenValue(n int)
func addTenPointer(n *int)

func resetValue(n int)
func resetPointer(n *int)

func swapValue(a, b int)
func swapPointer(a, b *int)

func appendValue(s []int)
func appendPointer(s *[]int)
```

Each pair has:

- One **value-based** function
- One **pointer-based** function

Only one choice per task will produce the correct output.

---

## Your Task

You will be given a `main` function with instructions in comments.

Your job is to **call the correct function** for each task.

You may:

- Call functions
- Pass values or addresses

You may **not**:

- Modify the functions
- Change function signatures
- Add new helper functions

---

## Tasks (Behavior Is the Test)

```go
func main() {

	// Task 1:
	// Make x become 15
	x := 5
	// CALL THE CORRECT FUNCTION
	fmt.Println(x) // expect: 15

	// Task 2:
	// Make y become 0
	y := 42
	// CALL THE CORRECT FUNCTION
	fmt.Println(y) // expect: 0

	// Task 3:
	// Swap a and b so output is: 9 3
	a, b := 3, 9
	// CALL THE CORRECT FUNCTION
	fmt.Println(a, b) // expect: 9 3

	// Task 4:
	// Make slice become [1 2 3 100]
	s := []int{1, 2, 3}
	// CALL THE CORRECT FUNCTION
	fmt.Println(s) // expect: [1 2 3 100]
}
```

---

## How This Tests Pointer Understanding

### Task 1 & 2 — Scalars

- Integers are copied
- Value-based functions cannot mutate caller state
- Only pointer-based functions work

If you do not understand `&` and `*`, these fail.

---

### Task 3 — Swapping Values

This is a classic trap.

```go
func swapValue(a, b int)
```

Cannot work because:

- `a` and `b` are copies

Only swapping through addresses mutates original variables.

---

### Task 4 — Slices (Subtle but Critical)

Slices contain:

- A pointer to an array
- A length
- A capacity

Passing a slice **copies the slice header**.

Appending may allocate a new array, breaking mutation.

Only a pointer to the slice guarantees length mutation.

This task separates surface-level understanding from real understanding.

---

## Requirements

1. Choose the correct function for each task
2. Use `&` only when mutation is required
3. Do not panic
4. Do not modify provided functions
5. Output must match expectations exactly

---

## Common Beginner Mistakes (Tests Will Catch)

- Calling value-based functions expecting mutation
- Forgetting to pass addresses
- Assuming slices always behave like references
- Confusing `*` (dereference) with `&` (address-of)
- Believing compilation success means correctness

---

## Key Ideas (Beginner-Friendly)

- Go passes values by default
- Pointers allow shared access to memory
- `&` gives an address
- `*` accesses the value at an address
- APIs using `*T` signal mutation
- Choosing the wrong function is a logic bug, not a syntax error

---

## Run Tests

```bash
go test ./quests/8.pointers -v
```
