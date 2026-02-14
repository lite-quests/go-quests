# Regular Expressions

## Concept

Regular expressions (regex) in Go are handled by the `regexp` package, which provides pattern matching and text manipulation capabilities. Regex patterns use special characters and sequences to define search criteria, making it easy to validate formats, extract data, and match complex string patterns. Understanding regex is essential for input validation, parsing, and text processing tasks.

## References

- https://gobyexample.com/regular-expressions
- https://pkg.go.dev/regexp
- https://github.com/google/re2/wiki/Syntax

## Quest

### Objective

Create a program that implements five validation functions using regular expressions. Each function checks if a string matches a specific pattern. You'll need to explore the references to construct the correct regex patterns.

### Requirements

#### Implement these five functions

```go
// IsOnlyNumbers returns true if the string contains only digits (0-9)
// Examples: "123" -> true, "12a" -> false, "" -> false
func IsOnlyNumbers(s string) bool

// IsOnlyAlphabets returns true if the string contains only letters (a-z, A-Z)
// Examples: "Hello" -> true, "Hello123" -> false, "" -> false
func IsOnlyAlphabets(s string) bool

// IsValidEmail returns true if the string is a valid email format
// Simple pattern: username@domain.extension
// Examples: "user@example.com" -> true, "invalid.email" -> false
func IsValidEmail(s string) bool

// ContainsGoQuest returns true if the string contains "go-quest" (case-insensitive)
// Examples: "learning go-quest" -> true, "Go-Quest rocks" -> true, "golang" -> false
func ContainsGoQuest(s string) bool

// IsValidUsername returns true if username is 3-16 characters, alphanumeric plus underscores
// Must start with a letter
// Examples: "user_123" -> true, "a" -> false, "123user" -> false
func IsValidUsername(s string) bool
```

#### Test cases to verify

Your functions should handle these test cases correctly:

```go
// IsOnlyNumbers
IsOnlyNumbers("12345")        // true
IsOnlyNumbers("123abc")       // false
IsOnlyNumbers("")             // false

// IsOnlyAlphabets
IsOnlyAlphabets("Hello")      // true
IsOnlyAlphabets("Hello123")   // false
IsOnlyAlphabets("")           // false

// IsValidEmail
IsValidEmail("user@example.com")     // true
IsValidEmail("test.user@domain.co")  // true
IsValidEmail("invalid.email")        // false
IsValidEmail("@example.com")         // false

// ContainsGoQuest
ContainsGoQuest("learning go-quest")      // true
ContainsGoQuest("Go-Quest is awesome")    // true
ContainsGoQuest("GO-QUEST")               // true
ContainsGoQuest("golang")                 // false

// IsValidUsername
IsValidUsername("user_123")    // true
IsValidUsername("JohnDoe")     // true
IsValidUsername("ab")          // false (too short)
IsValidUsername("123user")     // false (starts with number)
IsValidUsername("this_is_a_very_long_username")  // false (too long)
```

### Inputs

- Each function receives a single string parameter
- No external inputs required

### Outputs

- Each function returns a boolean value (`true` or `false`)

### Pattern Hints

Study the references to learn about:

- `^` and `$` - Start and end anchors
- `[0-9]` or `\d` - Digit patterns
- `[a-zA-Z]` - Letter patterns
- `+` - One or more occurrences
- `{n,m}` - Specific quantity ranges
- `(?i)` - Case-insensitive matching
- `.` - Any character
- `*` - Zero or more occurrences

**Why Regular Expressions Matter:**

- **Input validation**: Ensure user data matches expected formats before processing
- **Security**: Prevent invalid data from entering your system
- **Text processing**: Extract or find specific patterns in large text bodies
- **Data cleaning**: Identify and remove unwanted characters or formats
- **User experience**: Provide immediate feedback on form inputs

Mastering regex patterns is a valuable skill across all programming languages.

## Testing

To run the tests:

```bash
go test -v ./quests/023.regex
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestIsOnlyNumbers
--- PASS: TestIsOnlyNumbers (0.00s)
=== RUN   TestIsOnlyAlphabets
--- PASS: TestIsOnlyAlphabets (0.00s)
=== RUN   TestIsValidEmail
--- PASS: TestIsValidEmail (0.00s)
=== RUN   TestContainsGoQuest
--- PASS: TestContainsGoQuest (0.00s)
=== RUN   TestIsValidUsername
--- PASS: TestIsValidUsername (0.00s)
PASS
```
