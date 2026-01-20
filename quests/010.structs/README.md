# Go Structs & Methods

## Concept
Structs are typed collections of fields. Methods can be defined on structs to encapsulate behavior. Receivers (value vs. pointer) determine whether a method modifies the original struct or operates on a copy.

## References
- https://gobyexample.com/structs
- https://gobyexample.com/methods
- https://go.dev/tour/moretypes/15

## Quest

### Objective
Implement a `User` struct and associated methods to model a user profile, demonstrating correct usage of value and pointer receivers.

### Requirements

#### `NewUser`
- Function: `NewUser(id int, name, email string, age int) *User`
- Initialize and return a new `User` pointer with fields set from parameters.
- Must not return `nil`.

#### `IsAdult`
- Method: `(u *User) IsAdult() bool`
- Return `true` if `Age >= 18`, else `false`.
- Must not modify the user.

#### `DisplayName`
- Method: `(u User) DisplayName() string`
- Return formatted string: `"Name <Email>"`.
- Use **value receiver** (operates on a copy).

#### `UpdateEmail`
- Method: `(u *User) UpdateEmail(email string)`
- Update the `Email` field of the user.
- Use **pointer receiver** (mutates original).

#### `Birthday`
- Method: `(u *User) Birthday()`
- Increment `Age` by 1.
- Use **pointer receiver** (mutates original).

#### `CloneUser`
- Function: `CloneUser(u User) User`
- Return a new `User` instance with identical field values.
- Must return a copy, not a pointer.

### Inputs
- Varies by function/method (see above).

### Outputs
- Varies by function/method (see above).

### Examples
- `NewUser(1, "A", "e", 20)` → `&User{1, "A", "e", 20}`
- `u.DisplayName()` → `"A <e>"`
- `u.UpdateEmail("new")` → `Email` becomes `"new"`
- `u.Birthday()` → `Age` becomes `21`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/010.structs
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestNewUser
--- PASS: TestNewUser (0.00s)
=== RUN   TestUserIsAdult
--- PASS: TestUserIsAdult (0.00s)
=== RUN   TestDisplayName
--- PASS: TestDisplayName (0.00s)
=== RUN   TestUpdateEmail
--- PASS: TestUpdateEmail (0.00s)
=== RUN   TestBirthday
--- PASS: TestBirthday (0.00s)
=== RUN   TestCloneUser
--- PASS: TestCloneUser (0.00s)
PASS
```
