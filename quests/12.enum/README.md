# Go Enums Quest — Order Lifecycle State Machine

Your task is to implement an **order lifecycle system** using **enum-style constants** in Go.

This quest focuses on:

- Defining enum-like types using `iota`
- Modeling **finite state machines**
- Implementing `fmt.Stringer`
- Writing safe, explicit state transitions
- Handling invalid states defensively
- Understanding type safety with custom types

---

## Reference

- [https://gobyexample.com/enums](https://gobyexample.com/enums)
- [https://go.dev/ref/spec#Constants](https://go.dev/ref/spec#Constants)
- [https://go.dev/blog/constants](https://go.dev/blog/constants)
- [https://pkg.go.dev/fmt#Stringer](https://pkg.go.dev/fmt#Stringer)

---

## Core Idea

You are modeling the **lifecycle of an order** in an e-commerce system.

An order can only be in **one valid state at a time**, and **not all transitions are allowed**.

This is exactly what enums are good at.

---

## Enum Definition

You are given an enum-like type `OrderState` with an underlying `int`.

```go
type OrderState int
```

---

## Required States

Define the following states using `iota` **in this exact order**:

1. `StateCreated`
2. `StatePaid`
3. `StatePacked`
4. `StateShipped`
5. `StateDelivered`
6. `StateCancelled`
7. `StateReturned`

---

## String Representation

You must implement `fmt.Stringer` so that states print correctly.

### Required Mapping

| State          | String Value  |
| -------------- | ------------- |
| StateCreated   | `"created"`   |
| StatePaid      | `"paid"`      |
| StatePacked    | `"packed"`    |
| StateShipped   | `"shipped"`   |
| StateDelivered | `"delivered"` |
| StateCancelled | `"cancelled"` |
| StateReturned  | `"returned"`  |

### Constraint

- Use a `map[OrderState]string`
- Do **not** use a `switch` inside `String()`

```go
func (s OrderState) String() string
```

---

## State Transition Function

```go
func NextState(current OrderState, action string) OrderState
```

This function determines the **next state** of an order based on the current state and an action.

---

## Valid Actions

| Action      | Meaning                    |
| ----------- | -------------------------- |
| `"pay"`     | Customer completes payment |
| `"pack"`    | Seller packs the order     |
| `"ship"`    | Order is shipped           |
| `"deliver"` | Order is delivered         |
| `"cancel"`  | Order is cancelled         |
| `"return"`  | Customer returns the order |

---

## Transition Rules

### From `StateCreated`

- `"pay"` → `StatePaid`
- `"cancel"` → `StateCancelled`

---

### From `StatePaid`

- `"pack"` → `StatePacked`
- `"cancel"` → `StateCancelled`

---

### From `StatePacked`

- `"ship"` → `StateShipped`

---

### From `StateShipped`

- `"deliver"` → `StateDelivered`
- `"return"` → `StateReturned`

---

### From `StateDelivered`

- `"return"` → `StateReturned`

---

### Terminal States

The following states are **terminal** and cannot transition further:

- `StateCancelled`
- `StateReturned`

Any action on these states should **keep the state unchanged**.

---

## Invalid Transitions

- If an action is invalid for the current state:

  - Return the **current state unchanged**

- If the state itself is unknown:

  - **Panic with a descriptive error**

Example:

```go
panic(fmt.Errorf("unknown order state: %d", current))
```

---

## Type Safety Requirement

- `NextState` must accept **only** `OrderState`
- Passing a plain `int` must not compile

This is the primary benefit of enums in Go.

---

## Example Usage

```go
state := StateCreated
state = NextState(state, "pay")
fmt.Println(state) // paid

state = NextState(state, "pack")
fmt.Println(state) // packed

state = NextState(state, "ship")
fmt.Println(state) // shipped

state = NextState(state, "deliver")
fmt.Println(state) // delivered
```

---

## Requirements

1. Use `iota` correctly
2. Implement `fmt.Stringer`
3. Use a `switch` for transitions
4. No global mutation
5. No reflection
6. No string comparisons for states
7. Handle unknown states explicitly

---

## Common Beginner Mistakes (Tests Will Catch These)

- Forgetting to handle terminal states
- Returning zero-value on invalid transitions
- Using raw integers instead of `OrderState`
- Implementing `String()` with a switch
- Allowing illegal transitions
- Ignoring unknown enum values

---

## Key Ideas (What This Quest Teaches)

- Enums model **closed sets of states**
- `iota` produces predictable, ordered constants
- Custom types give compile-time safety
- State machines make illegal states unrepresentable
- `Stringer` improves debuggability and logs
- Explicit transitions prevent accidental behavior

---

## Mental Model

Enums are not “just constants”.

They are:

> **A contract that limits what values are allowed and how the system can evolve.**

---

## Run Tests

```bash
go test ./quests/12.enum -v
```
