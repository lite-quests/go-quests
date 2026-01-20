# Go Goroutines & Channels

## Concept
Goroutines are lightweight threads of execution. Channels are typed conduits for safe communication and synchronization between goroutines. Together, they enable concurrent processing without explicit locks or condition variables.

## References
- https://gobyexample.com/goroutines
- https://gobyexample.com/channels
- https://go.dev/tour/concurrency/1

## Quest

### Objective
Implement a basic concurrent request-response flow where a "Client" sends a request to a "Server" running in a separate goroutine, and waits for the response via a channel.

### Requirements

#### `SendRequest` (Client)
- Function: `SendRequest(input string) string`
- Create an unbuffered string channel.
- Start the `Server` function in a new **goroutine** (`go Server(...)`), passing the input and the channel.
- **Block** and wait for the response from the channel.
- Return the received response.

#### `Server` (Server)
- Function: `Server(request string, response chan string)`
- Runs concurrently (invoked with `go`).
- Format the output string as `"processed: <request>"`.
- **Send** the formatted string into the `response` channel.
- Do not close the channel.

### Inputs
- `input`: The raw request string (e.g., "ping").

### Outputs
- Returns the processed string (e.g., "processed: ping").

### Examples
- `SendRequest("hello")`
    1. Channel created.
    2. `go Server("hello", ch)` starts.
    3. `SendRequest` blocks on `<-ch`.
    4. `Server` sends `"processed: hello"`.
    5. `SendRequest` unblocks and returns `"processed: hello"`.

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/015.go_routine
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestSendRequest
--- PASS: TestSendRequest (0.00s)
PASS
```
