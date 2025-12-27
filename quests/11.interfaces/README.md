# Go Interfaces & Methods Quest — Polymorphic Payment System

Your task is to implement a **payment processing system** using Go **interfaces and methods**.

This quest focuses on:

- Defining interfaces (contracts)
- Implementing interfaces with structs
- Method receivers and method sets
- Interface-based polymorphism
- Runtime type inspection (type assertions)
- Designing behavior-first APIs

---

## Reference

- [https://gobyexample.com/methods](https://gobyexample.com/methods)
- [https://gobyexample.com/interfaces](https://gobyexample.com/interfaces)
- [https://go.dev/doc/effective_go#interfaces](https://go.dev/doc/effective_go#interfaces)
- [https://go.dev/tour/methods/9](https://go.dev/tour/methods/9)

---

## Core Idea

You are modeling **different payment methods** (e.g., Card, UPI, Crypto), all of which behave differently but must conform to a **single master interface**.

The system should not care _how_ payment is processed — only that it **can be processed**.

This is the essence of interfaces in Go.

---

## Master Interface

```go
type PaymentMethod interface {
	Process(amount float64) bool
	Provider() string
}
```

### Interface Meaning

Any type that implements **both**:

- `Process(amount float64) bool`
- `Provider() string`

**automatically satisfies** `PaymentMethod`.

No `implements` keyword.
No explicit declaration.
Pure behavior-based typing.

---

## Struct Implementations

You must implement **three different payment types**, each with **different internal state** and behavior.

---

### 1. `CardPayment`

```go
type CardPayment struct {
	CardNumber string
	Limit      float64
}
```

#### Required Methods

```go
func (c *CardPayment) Process(amount float64) bool
func (c CardPayment) Provider() string
```

#### Behavior

- `Process`

  - Returns `false` if `amount > Limit`
  - Deducts amount from `Limit` if successful
  - **Must mutate state**

- `Provider`

  - Returns `"CARD"`

Why pointer receiver?
Because processing a card **changes remaining credit**.

---

### 2. `UPIPayment`

```go
type UPIPayment struct {
	VPA string
}
```

#### Required Methods

```go
func (u UPIPayment) Process(amount float64) bool
func (u UPIPayment) Provider() string
```

#### Behavior

- Always returns `true` for `Process`
- Does **not** mutate state
- Provider returns `"UPI"`

Why value receiver?
Because UPI payments are stateless for this model.

---

### 3. `CryptoPayment`

```go
type CryptoPayment struct {
	Wallet  string
	Balance float64
}
```

#### Required Methods

```go
func (c *CryptoPayment) Process(amount float64) bool
func (c CryptoPayment) Provider() string
```

#### Behavior

- Fails if `amount > Balance`
- Deducts from balance on success
- Provider returns `"CRYPTO"`

---

## Polymorphic Function

```go
func Checkout(p PaymentMethod, amount float64) string
```

### What to Do

- Call `Process(amount)`
- If payment succeeds:

  - Return `"Payment successful via <provider>"`

- If it fails:

  - Return `"Payment failed via <provider>"`

### Critical Rule

`Checkout` **must not** know or care about concrete types.

No `if card`, `if upi`, `if crypto`.

Only interface methods allowed.

---

## Runtime Type Detection (Advanced)

```go
func DetectCrypto(p PaymentMethod) bool
```

### What to Do

- Return `true` if `p` is a `CryptoPayment`
- Otherwise return `false`

### Constraint

- Must use **type assertion**
- Must not panic

Example:

```go
if DetectCrypto(p) {
	// crypto-specific logic
}
```

---

## Example Usage

```go
card := &CardPayment{
	CardNumber: "1234",
	Limit:      1000,
}

upi := UPIPayment{
	VPA: "mani@upi",
}

crypto := &CryptoPayment{
	Wallet:  "0xabc",
	Balance: 5,
}

Checkout(card, 200)
Checkout(upi, 500)
Checkout(crypto, 3)
```

All of these must compile and work without changing `Checkout`.

---

## Requirements

1. Interface must define behavior, not data
2. Structs must implement interface **implicitly**
3. Use pointer receivers only where mutation is required
4. Do not use type switches inside `Checkout`
5. No global state
6. No reflection

---

## Common Beginner Mistakes (Tests Will Catch These)

- Using pointer receivers unnecessarily
- Expecting structs to “declare” interface implementation
- Mutating state with value receivers
- Writing type-specific logic in polymorphic functions
- Confusing interface values with concrete values

---

## Key Ideas (What This Quest Teaches)

- Interfaces define **capabilities**, not identities
- Method sets determine interface satisfaction
- Pointer vs value receivers affect interface compatibility
- Polymorphism emerges naturally from behavior
- Interfaces decouple _what_ from _how_

---

## Mental Model

> “If it walks like a duck and quacks like a duck — it _is_ a duck.”

In Go:

> “If it has the methods — it _is_ the interface.”

---

## Run Tests

```bash
go test ./quests/11.interfaces -v
```
