# Rate Limiting

## Concept

Rate limiting controls how frequently operations can occur. In Go, this is commonly implemented using channels and timers. A rate limiter allows a certain number of requests immediately (burst), then enforces a maximum rate for subsequent requests.

## References

- https://gobyexample.com/rate-limiting
- https://gobyexample.com/tickers
- https://go.dev/blog/pipelines

## Quest

### Objective

Implement `PingPong` to handle requests with rate limiting. Allow the first 4 requests to process immediately (burst capacity), then limit subsequent requests to 10 per second (one every 100ms). The function responds to "Ping" with "Pong" and vice versa, including timestamps.

### Requirements

#### `PingPong`

- Accept a slice of strings (requests) as input.
- Process the **first 4 requests immediately** without delay.
- After the initial burst, process **one request every 100ms** (10 requests/second).
- For each request:
  - `"Ping"` → print: `Pong HH:MM:SS.mmm`
  - `"Pong"` → print: `Ping HH:MM:SS.mmm`
  - Use timestamp format: `"15:04:05.000"`

#### Rate Limiting Strategy

- Use a **buffered channel** for burst capacity.
- Pre-fill the channel to allow instant processing.
- Use `time.Tick` or `time.NewTicker` to continuously supply tokens.
- Each request must acquire a token before processing.

### Inputs

- `requests []string`: Slice of "Ping" or "Pong" strings

### Outputs

- Prints response and timestamp for each request

### Example

```go
requests := []string{"Ping", "Pong", "Ping", "Ping", "Pong", "Ping", "Pong"}
PingPong(requests)
```

Output:

```text
Pong 14:32:01.123
Ping 14:32:01.123
Pong 14:32:01.123
Pong 14:32:01.123
Ping 14:32:01.224
Pong 14:32:01.324
Ping 14:32:01.424
```

**Behavior:**

- Requests 1-4: Same/similar timestamp (burst)
- Requests 5+: ~100ms intervals (rate limited)

## Testing

To run the tests:

```bash
go test -v ./quests/018.rate_limiting
```
