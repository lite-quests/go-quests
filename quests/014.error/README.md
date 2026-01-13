# Go Errors

## Concept
Errors in Go are values. Effective error handling involves using simple errors, sentinel errors (package-level variables), custom error types for structured data, and error wrapping to preserve context. Use `errors.Is` and `errors.As` to inspect errors.

## References
- https://go.dev/blog/go1.13-errors
- https://gobyexample.com/errors
- https://pkg.go.dev/errors

## Quest

### Objective
Implement a file validation system using sentinel errors, custom error types, error wrapping, and error inspection.

### Requirements

#### Variables & Types
- **Sentinel Errors** (Package-level):
    - `ErrEmptyFilename`: `"filename cannot be empty"`
    - `ErrFileTooLarge`: `"file size exceeds limit"`
- **Custom Error**: `ValidationError` struct with fields `Filename` (string) and `Reason` (string).
    - Implement `Error() string`: Return `"validation failed for '<Filename>': <Reason>"`.

#### Functions

1.  **`ValidateFilename(filename string) error`**
    -   Return `ErrEmptyFilename` if `filename` is empty.
    -   Return `nil` otherwise.

2.  **`ValidateFileSize(size int64, maxSize int64) error`**
    -   Return `ErrFileTooLarge` if `size > maxSize`.
    -   Return new error `"file size cannot be negative"` if `size < 0`.
    -   Return `nil` otherwise.

3.  **`ValidateFileExtension(filename string, allowedExts []string) error`**
    -   Check if `filename` ends with any string in `allowedExts`.
    -   If valid, return `nil`.
    -   If invalid, return a `*ValidationError` with `Filename` set and `Reason` set to `"unsupported file extension"`.

4.  **`ValidateFile(filename string, size int64, maxSize int64) error`**
    -   Call `ValidateFilename`. If error, wrap it with `"file validation failed: %w"`.
    -   Call `ValidateFileSize`. If error, wrap it with `"file validation failed: %w"`.
    -   Return `nil` if success.

5.  **`CanRetry(err error) bool`**
    -   Return `true` if `err` wraps a `*ValidationError` (unsupported extension can be fixed).
    -   Return `false` if `err` is `nil`, or wraps `ErrEmptyFilename`, `ErrFileTooLarge`, or any other error.

### Inputs
- Varies by function.

### Outputs
- Errors (Sentinel, Custom, Wrapped, or simple) and Booleans.

### Examples
- `ValidateFilename("")` → `ErrEmptyFilename`
- `ValidateFileSize(200, 100)` → `ErrFileTooLarge`
- `ValidateFileExtension("doc.exe", [".txt"])` → `&ValidationError{"doc.exe", "unsupported file extension"}`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/14.error
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestValidateFilename
--- PASS: TestValidateFilename (0.00s)
=== RUN   TestValidateFileSize
--- PASS: TestValidateFileSize (0.00s)
=== RUN   TestValidateFileExtension
--- PASS: TestValidateFileExtension (0.00s)
=== RUN   TestValidateFileWrapping
--- PASS: TestValidateFileWrapping (0.00s)
=== RUN   TestCanRetry
--- PASS: TestCanRetry (0.00s)
PASS
```
