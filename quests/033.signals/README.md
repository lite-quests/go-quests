# Signals: The Process Supervisor

## Concept

When you run a background server or worker (a "daemon") in production, the operating system communicates with it using Signals. In modern Go, graceful shutdown is typically handled via Contexts, while other control signals (like reloading configuration) are handled manually.

Additionally, many Go programs act as **supervisors** that spin up background child processes. When the supervisor is told to shut down, it should politely tell its child processes to shut down securely, by explicitly sending a term signal, rather than forcefully killing them.

- **`signal.NotifyContext(ctx, os.Interrupt)`**: Modern Go way to handle `SIGINT` (Ctrl+C). It returns a new Context that is automatically cancelled when `os.Interrupt` is received.
- **`signal.Notify(c, syscall.SIGHUP, ...)`**: The classic way to subscribe to non-termination signals.
- **`cmd.Process.Signal(syscall.SIGTERM)`**: Used to send a signal to a running child process.

## Quest

### Objective

You are building the master control loop for the Multiplayer Game Server itself (following the initialization step from the previous quest). The server acts as a supervisor that spins up a background database backup worker (`workerCmd`). Your server must handle `SIGINT` gracefully to shut down the worker, and handle `SIGHUP` and `SIGUSR1` for manual administrative commands.

### Requirements

Implement the function:
`func RunGameServer(ready chan bool, workerCmd string) error`

1. Use `signal.NotifyContext` with `context.Background()` to listen for `os.Interrupt` (`syscall.SIGINT`). This will be your graceful shutdown context. Defer calling its cancellation function (e.g., `defer stop()`).
2. Create a standard signal channel (e.g. `make(chan os.Signal, 1)`) and use `signal.Notify` to register it to receive `syscall.SIGHUP` and `syscall.SIGUSR1`.
3. Use `exec.Command` to prepare the child process using `workerCmd`. _Note: We use `Command` instead of `CommandContext` so we can manually send `SIGTERM` instead of letting Context brutally kill the process._
4. Start the child process in the background using `cmd.Start()`. If it fails to start, return the error.
5. Send `true` to the `ready` channel (`ready <- true`) to let the tests know you're running.
6. Enter an infinite `for { select { ... } }` loop to handle events:
   - When the signal channel receives `syscall.SIGHUP`: Print `"Reloading game rules..."` and continue the loop.
   - When the signal channel receives `syscall.SIGUSR1`: Print `"Dumping player coordinates..."` and continue the loop.
   - When the graceful shutdown context's `Done()` channel receives a value (`case <-ctx.Done():`):
     - Print `"Shutting down worker..."`
     - Send a `syscall.SIGTERM` to the child process using `cmd.Process.Signal(syscall.SIGTERM)`
     - Call `cmd.Wait()` to wait for the child process to finish.
     - Return the result of `cmd.Wait()` directly, even if it is an ExitError from the SIGTERM.

### Inputs

- `ready`: A `chan bool` that you must send `true` to once `cmd.Start()` has succeeded.
- `workerCmd`: The executable name or path of the worker command.

### Outputs

- `error`: Start errors or the final wait error, otherwise `nil`.

### Example

```go
// In tests, we will send signals to your server Process.
ready := make(chan bool)
go RunGameServer(ready, "sleep 100")
<-ready // Server running and child started!

// If someone runs `kill -HUP <pid>`:
// Output: Reloading game rules...

// If someone runs `kill -INT <pid>` (or presses Ctrl+C):
// Output: Shutting down worker...
// (And the function gracefully terminates the child and returns nil)
```

## Testing

```bash
go test -v ./quests/033.signals
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestRunGameServer
    solution_test.go:60: Process exited with wait error (acceptable if killed): signal: terminated
--- PASS: TestRunGameServer (0.21s)
PASS
```
