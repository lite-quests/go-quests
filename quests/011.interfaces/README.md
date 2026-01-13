# Go Interfaces

## Concept
Interfaces in Go define a set of method signatures (behavior). Any type that implements those methods explicitly satisfies the interface. This enables polymorphism, allowing different types to be treated uniformly based on what they *do* rather than what they *are*.

## References
- https://gobyexample.com/interfaces
- https://go.dev/tour/methods/9
- https://go.dev/blog/laws-of-reflection

## Quest

### Objective
Implement a polymorphic payment processing system with three distinct payment methods (`CardPayment`, `UPIPayment`, `CryptoPayment`) satisfying a common `PaymentMethod` interface.

### Requirements

#### Interface: `PaymentMethod`
- `Process(amount float64) bool`: Attempts to process payment, returning success status.
- `Provider() string`: Returns the provider name (e.g., "CARD").

#### Structs & Implementations

1.  **`CardPayment`**
    - Fields: `Limit` (float64).
    - `Process`: Return `true` and deduct `amount` from `Limit` if `amount <= Limit`. Else `false`. **(Pointer Receiver)**.
    - `Provider`: Return `"CARD"`.

2.  **`UPIPayment`**
    - Fields: `VPA` (string).
    - `Process`: Always return `true`. (No balance check). **(Value Receiver)**.
    - `Provider`: Return `"UPI"`.

3.  **`CryptoPayment`**
    - Fields: `Balance` (float64).
    - `Process`: Return `true` and deduct `amount` from `Balance` if `amount <= Balance`. Else `false`. **(Pointer Receiver)**.
    - `Provider`: Return `"CRYPTO"`.

#### Functions

-   **`Checkout(p PaymentMethod, amount float64) string`**
    -   Call `p.Process(amount)`.
    -   If success: Return `"Payment successful via <Provider>"`.
    -   If failure: Return `"Payment failed via <Provider>"`.

-   **`DetectCrypto(p PaymentMethod) bool`**
    -   Use **type assertion** to check if `p` is a `*CryptoPayment`.
    -   Return `true` if it is, `false` otherwise.

### Inputs
- Varies by struct and function (see above).

### Outputs
- Varies by function (boolean status, provider string, or status message).

### Examples
- `Checkout(&CardPayment{Limit: 100}, 50)` → `"Payment successful via CARD"`
- `Checkout(UPIPayment{VPA: "a@b"}, 500)` → `"Payment successful via UPI"`
- `DetectCrypto(&CardPayment{})` → `false`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/11.interfaces
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestPaymentMethodInterface
--- PASS: TestPaymentMethodInterface (0.00s)
=== RUN   TestCardPaymentProcess
--- PASS: TestCardPaymentProcess (0.00s)
=== RUN   TestUPIPaymentProcess
--- PASS: TestUPIPaymentProcess (0.00s)
=== RUN   TestCryptoPaymentProcess
--- PASS: TestCryptoPaymentProcess (0.00s)
=== RUN   TestProvider
--- PASS: TestProvider (0.00s)
=== RUN   TestCheckout
--- PASS: TestCheckout (0.00s)
=== RUN   TestDetectCrypto
--- PASS: TestDetectCrypto (0.00s)
PASS
```
