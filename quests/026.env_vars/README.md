# Environment Variables Config Loader

## Concept

Think of Environment Variables as the "settings menu" for your code. Just like you might change a game's difficulty from "Normal" to "Hard" without having to reinstall the entire game, environment variables let you change how your application connects to databases, enables debug modes, or handles secrets without having to rewrite or recompile the Go code.

Environment variables are the industry standard for configuring applications across different environments (local development vs. production). They are passed into the running program by the operating system.

In Go, the `os` package provides access to these variables:
- `os.Getenv("KEY")`: Returns the value as a string, or an empty string if not set.
- `os.LookupEnv("KEY")`: Returns the value and a boolean `(value, exists)`. Useful when you need to distinguish between an intentional empty value `""` vs a variable never being set.

**Crucial Note**: Environment variables from the OS are *always* returned as strings. If you need numbers or booleans (like a port number or debug flag), you must explicitly convert them using the `strconv` package.

## References
- [Go by Example: Environment Variables](https://gobyexample.com/environment-variables)
- [pkg.go.dev/os](https://pkg.go.dev/os)
- [12-Factor App: Config](https://12factor.net/config)

## Quest

### Objective
Implement a configuration loader that reads database settings and feature flags from environment variables.

### Requirements

1. Define a `Config` struct with the following fields:
   - `DBHost` (string)
   - `DBPort` (int)
   - `DebugMode` (bool)

2. Implement the function `LoadConfig() (Config, error)`.

3. The function must behave as follows:
   - **DBHost**: Read from `DB_HOST`. **Required**. Return an error if not set or empty.
   - **DBPort**: Read from `DB_PORT`. **Optional**. Default to `5432` if not set. Return an error if set but not a valid integer.
   - **DebugMode**: Read from `DEBUG_MODE`. **Optional**. Default to `false` if not set. Parse using `strconv.ParseBool`. Return an error if set but invalid.

### Inputs
- None (reads from system environment).

### Outputs
- `Config`: The populated struct.
- `error`: Non-nil if any validation fails.

### Example

```go
// Env: DB_HOST=localhost, DB_PORT=8080
config, err := LoadConfig()
// config.DBHost == "localhost"
// config.DBPort == 8080
// config.DebugMode == false
```

## Tips
- Use `os.Setenv` in tests (or your main function) to simulate values.
- converting string to int: `strconv.Atoi(s)`
- converting string to bool: `strconv.ParseBool(s)`

## Testing

```bash
go test -v ./quests/0003.env_vars
```
