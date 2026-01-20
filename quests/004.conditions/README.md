# Go Conditions

## Concept
Conditional logic controls the flow of execution based on boolean states. Go supports standard `if/else` structures for sequential logic and `switch` statements for multi-branch, discrete value evaluation or cleaner condition grouping.

## References
- https://gobyexample.com/if-else
- https://gobyexample.com/switch
- https://go.dev/tour/flowcontrol/5

## Quest

### Objective
Implement complex request classification using `if/else` and a grading system using `switch`.

### Requirements

#### `ClassifyRequest`
- Function: `ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string`
- Use `if`, `else if`, `else`. **Do not use switch**.
- Evaluate rules top-to-bottom. Return immediately upon match.
- **Rules**:
    1. **"INVALID"**: if `age <= 0` OR `balance < 0`.
    2. **"REJECTED"**: if `age < 18` OR `hasID` is false.
    3. **"VIP_ACCESS"**: if `isVIP` is true AND `balance >= 10000`.
    4. **"STANDARD_ACCESS"**: if `balance >= 1000`.
    5. **"LIMITED_ACCESS"**: All other cases.

#### `EvaluateGrade`
- Function: `EvaluateGrade(score int) string`
- Use `switch`. **Do not use if/else**.
- **Rules**:
    1. **"INVALID"**: if `score < 0` OR `score > 100`.
    2. **"A"**: 90–100.
    3. **"B"**: 80–89.
    4. **"C"**: 70–79.
    5. **"D"**: 60–69.
    6. **"F"**: 0–59.

### Inputs
- `ClassifyRequest`: `age` (int), `hasID` (bool), `balance` (float64), `isVIP` (bool)
- `EvaluateGrade`: `score` (int)

### Outputs
- Both return a `string` matching exactly one of the allowed return values defined in requirements (e.g. "VIP_ACCESS", "A").

### Examples
- `ClassifyRequest(30, true, 15000, true)` → `"VIP_ACCESS"`
- `ClassifyRequest(16, true, 100, false)` → `"REJECTED"`
- `EvaluateGrade(95)` → `"A"`
- `EvaluateGrade(105)` → `"INVALID"`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/004.conditions
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestClassifyRequest
--- PASS: TestClassifyRequest (0.00s)
=== RUN   TestEvaluateGrade
--- PASS: TestEvaluateGrade (0.00s)
PASS
```
