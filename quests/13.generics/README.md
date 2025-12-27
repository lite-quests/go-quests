# Go Generics Quest — Simple Generic Linked List

Your task is to implement a **singly-linked list** in Go using **generics**.

This quest focuses on:

- Defining **generic types** with type parameters
- Implementing **methods** on generic structs
- Using **slices to return list elements**
- Understanding **type-safe collections**
- Iterating over custom generic data structures

---

## Reference

- [https://gobyexample.com/generics](https://gobyexample.com/generics)
- [https://go.dev/doc/tutorial/generics](https://go.dev/doc/tutorial/generics)
- [https://pkg.go.dev/fmt](https://pkg.go.dev/fmt)

---

## Core Idea

You are implementing a **generic singly-linked list**.

- The list should store values of **any type** `T`.
- You will implement **methods** to manipulate and inspect the list.
- This helps understand **how generics work in Go**.

---

## Generic Type Definition

You need a `List[T]` struct and an internal `element[T]` struct:

```go
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}
```

---

## Methods to Implement

### 1. `Push`

```go
func (lst *List[T]) Push(v T)
```

- Adds a value to the **end** of the list
- Updates `head` and `tail` pointers

---

### 2. `Pop` (Optional)

```go
func (lst *List[T]) Pop() (T, bool)
```

- Removes and returns the **last element**
- Returns `false` if the list is empty

---

### 3. `AllElements`

```go
func (lst *List[T]) AllElements() []T
```

- Returns a **slice of all values** in the list
- Iterates from `head` to `tail`

---

## Example Usage

```go
lst := List[int]{}

lst.Push(10)
lst.Push(20)
lst.Push(30)

fmt.Println(lst.AllElements()) // [10 20 30]

val, ok := lst.Pop()
fmt.Println(val, ok)           // 30 true
fmt.Println(lst.AllElements()) // [10 20]

lst2 := List[string]{}
lst2.Push("foo")
lst2.Push("bar")
fmt.Println(lst2.AllElements()) // ["foo", "bar"]
```

---

## Requirements

1. Use **generic type parameters** correctly (`T any`)
2. Maintain **head and tail pointers** in methods
3. `AllElements` must return a **slice of `T`**
4. `Pop` should **return false** if the list is empty
5. Avoid using slices internally to store elements (except in `AllElements`)

---

## Common Beginner Mistakes (Tests Will Catch These)

- Forgetting to use `[T any]` in methods
- Returning `nil` slices instead of empty slices
- Not updating `tail` after `Push`
- `Pop` failing on single-element or empty list
- Using a slice internally to simulate the list instead of nodes

---

## Key Ideas (What This Quest Teaches)

- Generics allow **type-safe reusable data structures**
- Methods on generics require **type parameters to be repeated**
- Linked lists teach **pointers, iteration, and mutation**
- Returning slices from custom types is idiomatic in Go
- Maintaining internal invariants (`head` and `tail`) is essential

---

## Mental Model

> A generic linked list is like a conveyor belt:
> items can be added or removed, and it works for **any type**, but the internal nodes keep the structure consistent.

---

## Run Tests

```bash
go test ./quests/13.generics -v
```
