# Go Enums

## Concept
Go does not have a native `enum` type. Instead, enums are simulated using `const` groups with `iota` and a custom type. This pattern allows for type safety and easy integer-to-string conversion by implementing the `fmt.Stringer` interface.

## References
- https://gobyexample.com/enums
- https://pkg.go.dev/fmt#Stringer
- https://go.dev/blog/constants

## Quest

### Objective
Implement an order lifecycle state machine using `OrderState` enum. Define valid state strings and enforce transition rules.

### Requirements

#### `OrderState` Enum
- Define `OrderState` as an `int`.
- Use `iota` to define constants in this **exact order**:
  1. `StateCreated`
  2. `StatePaid`
  3. `StatePacked`
  4. `StateShipped`
  5. `StateDelivered`
  6. `StateCancelled`
  7. `StateReturned`

#### `String` Method
- Method: `(s OrderState) String() string`
- Return the string representation of the state using a map lookup (do not use switch):
    - `StateCreated` → `"created"`
    - `StatePaid` → `"paid"`
    - `StatePacked` → `"packed"`
    - `StateShipped` → `"shipped"`
    - `StateDelivered` → `"delivered"`
    - `StateCancelled` → `"cancelled"`
    - `StateReturned` → `"returned"`

#### `NextState`
- Function: `NextState(current OrderState, action string) OrderState`
- Implement state transitions based on valid actions:
    - `StateCreated` + `"pay"` → `StatePaid`
    - `StateCreated` + `"cancel"` → `StateCancelled`
    - `StatePaid` + `"pack"` → `StatePacked`
    - `StatePaid` + `"cancel"` → `StateCancelled`
    - `StatePacked` + `"ship"` → `StateShipped`
    - `StateShipped` + `"deliver"` → `StateDelivered`
    - `StateShipped` + `"return"` → `StateReturned`
    - `StateDelivered` + `"return"` → `StateReturned`
- **Invalid Actions**: Return `current` state unchanged.
- **Terminal States**: `StateCancelled` and `StateReturned` cannot transition.
- **Unknown States**: Panic with a descriptive error if `current` is not a valid state.

### Inputs
- `current`: The current `OrderState`.
- `action`: A string describing the event (e.g., `"pay"`, `"ship"`).

### Outputs
- Returns the new `OrderState`.

### Examples
- `NextState(StateCreated, "pay")` → `StatePaid`
- `NextState(StateCreated, "pack")` → `StateCreated` (Invalid transition)
- `NextState(StateCancelled, "pay")` → `StateCancelled` (Terminal state)
- `StatePaid.String()` → `"paid"`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/12.enum
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestStatesExist
--- PASS: TestStatesExist (0.00s)
=== RUN   TestOrderStateString
--- PASS: TestOrderStateString (0.00s)
=== RUN   TestValidTransitions
--- PASS: TestValidTransitions (0.00s)
=== RUN   TestTerminalStates
--- PASS: TestTerminalStates (0.00s)
=== RUN   TestUnknownStatePanics
--- PASS: TestUnknownStatePanics (0.00s)
PASS
```
