# Hello Go

## Concept
Standard output (stdout) is the primary output stream for command-line applications. In Go, the `fmt` package provides formatted I/O functions to write to stdout.

## References
- https://go.dev/tour/basics/1
- https://gobyexample.com/hello-world
- https://pkg.go.dev/fmt

## Quest

### Objective
Implement a function to print a specific string to standard output.

### Requirements
- Function: `HelloGo()`
- Package: `hello_world`
- The function must print exactly "Yo! Hello Go" to stdout.
- Do not add a newline character at the end unless implicitly added by the print function (check if `fmt.Print` or `fmt.Println` is required; tests check for exact bytes). Note: `fmt.Print` does not add a newline.

### Inputs
None

### Outputs
- Type: `string` (to stdout)
- Format: Plain text
- Expected behavior: "Yo! Hello Go" is written to stdout.

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/001.hello_go
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
PASS
```