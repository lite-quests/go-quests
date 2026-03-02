# Spawning Processes: Version Compatibility & Timeout Shield

## Concept

A very common use case for Go is building CLI tools (like Terraform, Docker CLI, or Kubernetes `kubectl`). Often, these tools act as orchestrators and require other software to be installed on the host machine.

More importantly, it's not enough to check if a tool exists; we often need to check its version. And because we don't control these third-party tools, we must ensure they don't hang and freeze our program.

The `os/exec` package provides the `Cmd` struct to execute external commands.

- **`exec.CommandContext(ctx, name, arg...)`**: Creates a `Cmd` representing the command to run, guarded by a Context. If the context times out, the child process is killed.
- **`cmd.Output()`**: Runs the command, waits for it to complete, and returns its standard output.

## Quest

### Objective

You are building the initialization step of a deployment CLI that validates dependencies before starting a Multiplayer Game Server. Before your tool runs, it needs to verify that a required third-party software (like `go`) is installed and is the correct version. To ensure your CLI doesn't hang if the tool malfunctions, you must also enforce a timeout using `exec.CommandContext`.

### Requirements

Implement the function:
`func CheckToolVersion(ctx context.Context, tool string, versionString string) error`

1. Use `exec.CommandContext` to prepare the execution of the `tool` with the argument `"version"`. (e.g., if the tool is `"go"`, the equivalent terminal command is `go version`). Make sure to pass the provided `ctx`!
2. Call `.Output()` on the command to execute it and capture its standard output.
3. If `.Output()` returns an error (which happens if the process times out, exits with a non-zero status, or the executable doesn't exist), return that error as-is, do not wrap it.
4. Convert the returned output byte slice into a string.
5. If the resulting string does not contain the `versionString`, return an error formatted exactly as:
   `fmt.Errorf("version mismatch: expected %s, got %s", versionString, outputString)`
   _(Note: The `outputString` should be the raw string returned by `.Output()`, stripped of any whitespace/newlines at the ends using `strings.TrimSpace` so the error message is clean)_.
6. If the output string contains the `versionString`, return `nil`.

### Inputs

- `ctx`: A `context.Context` to handle timeouts.
- `tool`: A string representing the tool to check (e.g., `"go"`).
- `versionString`: A substring to look for in the output (e.g., `"go1.20"`).

### Outputs

- `error`: The error returned by `.Output()`, or a version mismatch error, otherwise `nil`.

### Example

```go
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

err := CheckToolVersion(ctx, "go", "go1.")
// err: nil (assuming your Go version contains 'go1.')

err = CheckToolVersion(ctx, "go", "version-not-real")
// fmt.Println(err)
// Output: "version mismatch: expected version-not-real, got go version go..."
```

## Tips

- Check `strings.Contains` and `strings.TrimSpace` from the `"strings"` package.
- `cmd.Output()` returns `([]byte, error)`. `string(b)` converts `[]byte` to a string.

## Testing

```bash
go test -v ./quests/031.processes
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestCheckToolVersion
=== RUN   TestCheckToolVersion/success_version_match
=== RUN   TestCheckToolVersion/fail_version_mismatch
=== RUN   TestCheckToolVersion/fail_timeout
--- PASS: TestCheckToolVersion (0.53s)
    --- PASS: TestCheckToolVersion/success_version_match (0.01s)
    --- PASS: TestCheckToolVersion/fail_version_mismatch (0.01s)
    --- PASS: TestCheckToolVersion/fail_timeout (0.50s)
PASS
```
