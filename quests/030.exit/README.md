Here is the **final README.md** for your `os.Exit` quest, using the **exact same side headings and structure** as your goroutines example:

---

# Go Exit Status

## Concept

In Go, `os.Exit(code)` immediately terminates the current process with the given exit status code. Unlike a normal function return, `os.Exit` stops execution instantly and does not run any deferred functions.

Exit codes are used by command-line programs to signal success (`0`) or failure (`non-zero`). This allows the operating system and other programs to detect whether a program completed successfully.

Understanding `os.Exit` is important when building CLI tools and handling fatal errors.

## References

* [https://gobyexample.com/exit](https://gobyexample.com/exit)
* [https://pkg.go.dev/os#Exit](https://pkg.go.dev/os#Exit)
* [https://go.dev/doc/effective_go#errors](https://go.dev/doc/effective_go#errors)

## Quest

### Objective

Implement a function that terminates the program immediately with a specific exit status code.

### Requirements

#### `ExitWithStatus`

* Function: `ExitWithStatus()`

* Package: `exit_quest`

* Call `os.Exit(3)` to terminate the program.

* The deferred statement must NOT execute.

* The program must exit with status code `3`.

### Inputs

None

### Outputs

None (process exits immediately)

### Examples

Calling:

```go
ExitWithStatus()
```

Execution flow:

1. Deferred print is scheduled.
2. `os.Exit(3)` is called.
3. Program terminates immediately.
4. Deferred function is NOT executed.
5. Exit status is `3`.

### Testing

To run the tests, execute the following command from the root directory:

```bash
go test -v ./quests/016.exit_status
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestExitWithStatus
--- PASS: TestExitWithStatus (0.00s)
PASS
```
