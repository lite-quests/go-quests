# CLI File Processor

## Concept

Command-line interfaces (CLIs) are a primary way to interact with Go programs. Robust CLIs often support **subcommands** (like `git commit`, `docker run`) and **flags** (like `--verbose`, `-n`).

Go's standard library `os` provides access to raw arguments via `os.Args`, and the `flag` package helps parse them.

### `os.Args`

`os.Args` is a slice of strings representing the arguments passed to the program.

- `os.Args[0]` is the program name.
- `os.Args[1:]` are the user's arguments.

### Subcommands with `flag.NewFlagSet`

To implement subcommands, we inspect the first argument and then use a dedicated `flag.FlagSet` for that command.

```go
switch os.Args[1] {
case "foo":
    // Create a flag set for the "foo" command
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    fooFlag := fooCmd.Bool("enable", false, "desc")

    // Parse arguments specific to "foo" (skipping program name and subcommand)
    fooCmd.Parse(os.Args[2:])
    // ...
}
```

## References

- [Go by Example: Command-Line Arguments](https://gobyexample.com/command-line-arguments)
- [Go by Example: Command-Line Flags](https://gobyexample.com/command-line-flags)
- [Go by Example: Command-Line Subcommands](https://gobyexample.com/command-line-subcommands)
- [pkg.go.dev/flag](https://pkg.go.dev/flag)

## Quest

### Objective

Build a `file-cli` tool that can process text files in two ways: counting lines and searching for text.

### Requirements

Implement the function `RunCLI(args []string) error`.

The function should handle two subcommands: `count` and `search`.

#### 1. `count` Subcommand

Counts the number of lines in a file.

- **Usage**: `file-cli count <filename>`
- **Output**: Prints `lines: <number>` to stdout.
- **Error**: If file format is invalid or file missing, return an error.

#### 2. `search` Subcommand

Searches for lines containing a specific string pattern.

- **Usage**: `file-cli search --case-insensitive <pattern> <filename>`
- **Flags**:
  - `--case-insensitive` (bool): If set, case is ignored (e.g., "Foo" matches "foo"). Default is false.
- **Output**: Prints matching lines to stdout.
- **Error**: If arguments are missing or file cannot be read, return an error.

#### General Error Handling

- If no subcommand is provided (args length < 2), return an error.
- If an unknown subcommand is provided, return an error.
- If required arguments (like filename) are missing after flags, return an error.

### Inputs

- `args []string`: The full argument slice (including program name).

### Outputs

- **Return**: `error` if any operation fails.
- **Stdout**: The command results.

### Example

```go
args := []string{"./app", "count", "test.txt"}
RunCLI(args)
// Output: lines: 5
```

```go
args := []string{"./app", "search", "--case-insensitive", "hello", "test.txt"}
RunCLI(args)
// Output:
// Hello World
// hello there
```

## Tips

- Use `os.Open(filename)` to open a file.
- Use `bufio.NewScanner(file)` to read line by line.
- Remember `flag.Parse()` consumes flags. `flag.Args()` (or `flagSet.Args()`) returns the remaining non-flag arguments (like the filename).
- `flagSet.Parse` should be called with `args[2:]` because `args[0]` is program name and `args[1]` is the subcommand.

## Testing

To run the tests:

```bash
go test -v ./quests/1002.cli
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestRunCLI_Search/search_case-sensitive
=== RUN   TestRunCLI_Search/search_case-insensitive
--- PASS: TestRunCLI_Search (0.00s)
    --- PASS: TestRunCLI_Search/search_case-sensitive (0.00s)
    --- PASS: TestRunCLI_Search/search_case-insensitive (0.00s)
=== RUN   TestRunCLI_Errors
=== RUN   TestRunCLI_Errors/no_args
=== RUN   TestRunCLI_Errors/unknown_subcommand
=== RUN   TestRunCLI_Errors/missing_file_arg_for_count
=== RUN   TestRunCLI_Errors/missing_file_arg_for_serach
--- PASS: TestRunCLI_Errors (0.00s)
    --- PASS: TestRunCLI_Errors/no_args (0.00s)
    --- PASS: TestRunCLI_Errors/unknown_subcommand (0.00s)
    --- PASS: TestRunCLI_Errors/missing_file_arg_for_count (0.00s)
    --- PASS: TestRunCLI_Errors/missing_file_arg_for_serach (0.00s)
PASS
```
