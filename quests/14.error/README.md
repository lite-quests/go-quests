# Go Errors Quest — File Validator

Your task is to implement a **file validation system** using Go's error handling patterns.

This quest focuses on:

- Creating simple errors with `errors.New()`
- Wrapping errors with `fmt.Errorf` and `%w`
- Defining sentinel errors for common cases
- Implementing a basic custom error type
- Using `errors.Is()` and `errors.As()` for error checking

---

## Reference

- [https://go.dev/blog/error-handling-and-go](https://go.dev/blog/error-handling-and-go)
- [https://go.dev/blog/go1.13-errors](https://go.dev/blog/go1.13-errors)
- [https://pkg.go.dev/errors](https://pkg.go.dev/errors)

---

## The Scenario

You're building a file validator that checks if files meet certain requirements before processing them.

Files can fail validation for different reasons, and you need to handle each case appropriately.

---

## Part 1: Sentinel Errors

Define these **package-level variables**:

```go
var (
    ErrEmptyFilename = errors.New("filename cannot be empty")
    ErrFileTooLarge  = errors.New("file size exceeds limit")
)
```

**What are sentinel errors?**

- Predefined errors that represent specific conditions
- Declared at package level
- Callers can check for them using `errors.Is()`

---

## Part 2: Custom Error Type

Create a custom error type that includes the filename and size:

```go
type ValidationError struct {
    Filename string
    Size     int64
    Reason   string
}
```

Implement the `error` interface:

```go
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for '%s' (%d bytes): %s",
        e.Filename, e.Size, e.Reason)
}
```

**Important:** Use a **pointer receiver** (`*ValidationError`)

---

## Part 3: Validation Functions

### Function 1: `ValidateFilename`

```go
func ValidateFilename(filename string) error
```

**Rules:**

- If filename is empty (`""`), return `ErrEmptyFilename`
- Otherwise return `nil`

---

### Function 2: `ValidateFileSize`

```go
func ValidateFileSize(size int64, maxSize int64) error
```

**Rules:**

- If `size > maxSize`, return `ErrFileTooLarge`
- If `size < 0`, return a **simple error** using `errors.New("file size cannot be negative")`
- Otherwise return `nil`

**Note:** The negative size error is NOT a sentinel error - it's created on the spot.

---

### Function 3: `ValidateFileExtension`

```go
func ValidateFileExtension(filename string, allowedExts []string) error
```

**Rules:**

- Extract the file extension from `filename` (everything after the last `.`)
- If the extension is not in `allowedExts`, return a **ValidationError** with:
  - `Filename: filename`
  - `Size: 0`
  - `Reason: "unsupported file extension"`
- Otherwise return `nil`

**Hint:** Use `strings` package functions like `strings.LastIndex()` or `strings.HasSuffix()`

**Example:**

```go
ValidateFileExtension("doc.pdf", []string{".txt", ".md"})
// Returns ValidationError with reason "unsupported file extension"
```

---

### Function 4: `ValidateFile`

```go
func ValidateFile(filename string, size int64, maxSize int64) error
```

This function combines all validations:

1. Call `ValidateFilename(filename)`

   - If error, **wrap it** with: `"file validation failed: %w"`

2. Call `ValidateFileSize(size, maxSize)`

   - If error, **wrap it** with: `"file validation failed: %w"`

3. If all checks pass, return `nil`

**Key concept:** Wrapping adds context while preserving the original error.

---

## Part 4: Error Checking Function

```go
func CanRetry(err error) bool
```

Determine if validation can be retried after fixing the issue:

**Rules:**

- If `err` is `nil`, return `false`
- If `err` is (or wraps) `ErrFileTooLarge`, return `false` (can't fix size)
- If `err` is (or wraps) `ErrEmptyFilename`, return `false` (can't fix empty name)
- If `err` is a `*ValidationError`, return `true` (extension can be fixed)
- For any other error, return `false`

**Hints:**

- Use `errors.Is(err, ErrFileTooLarge)` to check sentinel errors
- Use `errors.As(err, &validationErr)` to check for custom error type

---

## Example Usage

```go
// Sentinel error
err := ValidateFilename("")
if errors.Is(err, ErrEmptyFilename) {
    fmt.Println("Please provide a filename")
}

// Wrapped error
err = ValidateFile("", 100, 1000)
if errors.Is(err, ErrEmptyFilename) {
    fmt.Println("Still caught the wrapped error!")
}

// Custom error
err = ValidateFileExtension("doc.pdf", []string{".txt", ".md"})
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    fmt.Printf("File: %s, Reason: %s\n",
        validationErr.Filename, validationErr.Reason)
}

// Retry logic
if CanRetry(err) {
    fmt.Println("User can fix this and retry")
}
```

---

## Run Tests

```bash
go test ./error/solution -v
```
