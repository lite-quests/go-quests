# Go Conditionals Quest

# If/Else Challenge

Your task is to implement a single function using **`if`, `else if`, and `else`** in Go.
This quest is designed to test **real conditional reasoning**, not just syntax.

## Reference

- [https://gobyexample.com/if-else](https://gobyexample.com/if-else)
- [https://gobyexample.com/booleans](https://gobyexample.com/booleans)

---

## Function: `ClassifyRequest`

```go
func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string
```

---

## Description

You are building a **request classification engine**.

Given user attributes (`age`, `hasID`, `balance`, `isVIP`), determine the **access level** based on a strict set of rules.

Exactly **one** classification string must be returned.

---

## Valid Return Values

The function must return **one of the following strings**:

- `"INVALID"`
- `"REJECTED"`
- `"VIP_ACCESS"`
- `"STANDARD_ACCESS"`
- `"LIMITED_ACCESS"`

No other values are allowed.

---

## Business Rules (Order Matters)

Rules must be evaluated **top to bottom**.
Once a condition matches, return immediately.

---

### 1. Invalid Request

Return `"INVALID"` if **any** of the following are true:

- `age <= 0`
- `balance < 0`

---

### 2. Rejected

Return `"REJECTED"` if **either** of the following are true:

- `age < 18`
- `hasID == false`

---

### 3. VIP Access

Return `"VIP_ACCESS"` if **both** are true:

- `isVIP == true`
- `balance >= 10000`

---

### 4. Standard Access

Return `"STANDARD_ACCESS"` if:

- `balance >= 1000`

---

### 5. Limited Access

All remaining **valid** requests fall into this category.

Return `"LIMITED_ACCESS"`.

---

## Requirements

1. Use `if / else if / else`
2. Use **compound conditions** (`&&`, `||`)
3. Use **early returns**
4. Do **not** use `switch`
5. Do **not** reorder the rules
6. Do **not** nest `if` blocks more than one level deep
7. Do not print anything
8. The function must be deterministic and side-effect free

---

## Examples

| age | hasID | balance | isVIP | Output          |
| --: | :---: | ------: | :---: | --------------- |
|  -1 | true  |     500 | false | INVALID         |
|  16 | true  |    2000 | false | REJECTED        |
|  25 | false |    5000 | true  | REJECTED        |
|  30 | true  |   15000 | true  | VIP_ACCESS      |
|  40 | true  |    3000 | false | STANDARD_ACCESS |
|  22 | true  |     200 | false | LIMITED_ACCESS  |

---

## Common Mistakes (Tests Will Catch These)

- Using multiple independent `if` statements
- Checking VIP before validating age/ID
- Forgetting early returns
- Allowing invalid inputs to pass through
- Returning default values incorrectly

---

## Key Ideas

- Condition **ordering is logic**
- `else if` enforces exclusivity
- Short-circuit logic (`&&`, `||`) matters
- Real-world validation is sequential, not parallel

---

# Switch Challenge

Your task is to implement a function using **`switch` statements** in Go.

This quest focuses on:

- Discrete states
- Grouped cases
- Clean control flow
- Avoiding unnecessary condition chains

---

## Reference

- [https://gobyexample.com/switch](https://gobyexample.com/switch)

---

## Function: `EvaluateGrade`

```go
func EvaluateGrade(score int) string
```

---

## Description

You are implementing a **grading engine**.

Given a numeric `score` between `0` and `100`, determine the corresponding **letter grade**.

---

## Valid Return Values

The function must return **exactly one** of the following:

- `"INVALID"`
- `"A"`
- `"B"`
- `"C"`
- `"D"`
- `"F"`

No other values are allowed.

---

## Rules (Strict)

Use a **switch statement** to evaluate the score.

---

### 1. Invalid Score

Return `"INVALID"` if:

- `score < 0`
- `score > 100`

---

### 2. Grade Mapping

| Score Range | Grade |
| ----------: | :---- |
|      90–100 | A     |
|       80–89 | B     |
|       70–79 | C     |
|       60–69 | D     |
|        0–59 | F     |

---

## Requirements

1. Use a `switch` statement
2. Do **not** use `if / else`
3. Do **not** use `fallthrough`
4. Use **case grouping** where appropriate
5. Ensure exactly one return path
6. Do not print anything

---

## Example Inputs

- `EvaluateGrade(95)` → `"A"`
- `EvaluateGrade(82)` → `"B"`
- `EvaluateGrade(70)` → `"C"`
- `EvaluateGrade(61)` → `"D"`
- `EvaluateGrade(40)` → `"F"`
- `EvaluateGrade(-5)` → `"INVALID"`
- `EvaluateGrade(120)` → `"INVALID"`

---

## Hint (Not Code)

You are expected to use a **boolean switch**:

```go
switch {
case condition:
    ...
}
```

This keeps the logic readable and idiomatic.

---

## Common Mistakes (Tests Will Catch These)

- Using chained `if / else`
- Forgetting bounds checking
- Misclassifying edge values (e.g. 89, 90)
- Using `fallthrough`
- Returning default values incorrectly

---

## Key Ideas

- `switch` expresses **discrete domains**
- Boolean switches replace long conditional chains cleanly
- Grouped cases improve readability
- `default` should represent the final valid state

---

## Run Tests

```bash
go test ./quests/4.conditions -v
```

Do not add logging or printing.
Only return the correct classification string.
