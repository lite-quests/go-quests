# Go Structs Quest — Immutable User Profile

Your task is to implement a **user profile system** using Go structs and methods.

This quest focuses on:

- Defining and initializing structs
- Value vs pointer receivers
- Updating struct fields safely
- Returning copies vs mutating state
- Encapsulation through methods
- Avoiding unintended mutations

---

## Reference

- [https://gobyexample.com/structs](https://gobyexample.com/structs)
- [https://gobyexample.com/methods](https://gobyexample.com/methods)
- [https://go.dev/tour/moretypes/2](https://go.dev/tour/moretypes/2)

---

## Struct

```go
type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}
```

---

## Functions & Methods to Implement

```go
func NewUser(id int, name string, age int, email string) User

func (u User) IsAdult() bool
func (u User) DisplayName() string

func (u *User) UpdateEmail(email string)
func (u *User) Birthday()

func CloneUser(u User) User
```

---

## Description

You are building a **simple user profile domain model**.

The struct represents a user’s state, and methods define valid interactions with that state.

Some methods must **mutate the user**, while others must be **read-only**.

Tests will ensure that you understand the difference.

---

## Function Behavior (Very Explicit)

---

### 1. `NewUser`

```go
func NewUser(id int, name string, age int, email string) User
```

What to do:

- Create a new `User` struct
- Assign all fields from parameters
- Return the struct **by value**

What NOT to do:

- Do not return a pointer
- Do not use global variables

Example:

```go
u := NewUser(1, "Mani", 20, "mani@example.com")
```

---

### 2. `IsAdult`

```go
func (u User) IsAdult() bool
```

What to return:

- `true` if `Age >= 18`
- `false` otherwise

Requirements:

- Use a **value receiver**
- Do not mutate the user

---

### 3. `DisplayName`

```go
func (u User) DisplayName() string
```

What to return:

- `"Name <Email>"`

Example:

```go
u.DisplayName()
// "Mani <mani@example.com>"
```

Requirements:

- Value receiver only
- Must not allocate unnecessary structs

---

### 4. `UpdateEmail`

```go
func (u *User) UpdateEmail(email string)
```

What to do:

- Update the user’s `Email` field
- Mutate the original user

Requirements:

- Must use a **pointer receiver**

Example:

```go
u.UpdateEmail("new@example.com")
```

---

### 5. `Birthday`

```go
func (u *User) Birthday()
```

What to do:

- Increment `Age` by 1

Important:

- This method **must mutate** the struct

---

### 6. `CloneUser`

```go
func CloneUser(u User) User
```

What to do:

- Return a **copy** of the given user

Why this exists:

- To make the difference between copying a struct and sharing a pointer explicit

Example:

```go
u1 := NewUser(1, "A", 20, "a@x.com")
u2 := CloneUser(u1)

u2.Birthday()
u1.Age // still 20
```

---

## Requirements

1. Use a struct, not a map
2. Use both value and pointer receivers correctly
3. Do not return pointers unless explicitly asked
4. No global state
5. No unnecessary allocations
6. Fields must be accessed through methods where applicable

---

## Common Beginner Mistakes (Tests Will Catch These)

- Using pointer receivers everywhere
- Mutating state in value-receiver methods
- Returning pointers from `NewUser`
- Forgetting to dereference in mutating methods
- Confusing copying with aliasing

---

## Key Ideas (What This Quest Teaches)

- Structs model **state + behavior**
- Value receivers protect immutability
- Pointer receivers enable mutation
- Copying structs is cheap and explicit
- Methods define a domain boundary

---

## Run Tests

```bash
go test ./quests/10.structs -v
```
