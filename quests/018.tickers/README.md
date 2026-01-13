# Go Tickers

## Concept
Tickers define a channel that delivers a "tick" of a clock at intervals. By using `select`, a goroutine can wait on multiple tickers and other channels (like a stop signal) simultaneously, allowing for coordinated periodic tasks.

## References
- https://gobyexample.com/tickers
- https://pkg.go.dev/time#Ticker
- https://go.dev/tour/concurrency/5

## Quest

### Objective
Implement `Ticker` to print messages at fixed intervals using two separate tickers, running for a specific duration before exiting cleanly.

### Requirements

#### Tickers
- **`helloTicker`**: Ticks every **500 ms**.
- **`worldTicker`**: Ticks every **1000 ms**.

#### Behavior
- Run a loop using `select` to handle ticker events.
- On `helloTicker` tick: Print `"hello"`.
- On `worldTicker` tick: Print `"world"`.
- **Duration**: The loop must run for exactly **3 seconds**.
- After 3 seconds:
    - Stop both tickers.
    - Exit the function (and goroutine).

#### Output Format
- Plain strings followed by a newline (handled by `fmt.Println` or equivalent).
- No extra text.

### Inputs
- None.

### Outputs
- Sequence of "hello" and "world" printed to stdout.

### Examples
Expected sequence over 3 seconds:
```text
hello
hello
world
hello
hello
world
hello
hello
world
```
(Note: Exact order of simultaneous ticks depends on runtime scheduling, but the tests verify this specific sequence).

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/18.tickers
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestTicker
--- PASS: TestTicker (3.00s)
PASS
```
