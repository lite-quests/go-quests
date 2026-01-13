# Go Channels

## Concept
Channels act as typed conduits for data. Correct usage requires understanding blocking behavior: sending blocks until a receiver is ready, and receiving blocks until a sender is ready. Functions must strictly adhere to their designated roles (producer, consumer, transformer) to avoid deadlocks.

## References
- https://gobyexample.com/channels
- https://gobyexample.com/channel-directions
- https://go.dev/tour/concurrency/2

## Quest

### Objective
Implement three channel manipulation functions observing strict read/write responsibilities and preventing deadlocks.

### Requirements

#### `ReadFromBoth`
- Function: `ReadFromBoth(ch1 chan string, ch2 chan string) string`
- **Receive** one value from `ch1`.
- **Receive** one value from `ch2`.
- Return string formatted as: `"read: <val1> & <val2>"`.
- Blocking behavior is expected.

#### `WriteToBoth`
- Function: `WriteToBoth(ch1 chan string, ch2 chan string, msg string)`
- **Send** formatted string `"write: <msg>"` to `ch1`.
- **Send** formatted string `"write: <msg>"` to `ch2`.
- **Constraint**: This function must **not block** the caller (the tests call it synchronously). Use a goroutine to perform the sends.

#### `ReadThenWrite`
- Function: `ReadThenWrite(input chan string, output chan string)`
- **Receive** one value from `input`.
- **Send** formatted string `"transform: <val>"` to `output`.
- Blocking behavior is expected.

### Inputs
- Channels of type `string` and plain `string` messages.

### Outputs
- Formatted strings returned or sent to channels.

### Examples
- `ReadFromBoth(ch1, ch2)` (where ch1="A", ch2="B") → `"read: A & B"`
- `WriteToBoth(ch1, ch2, "hi")` → `ch1` receives `"write: hi"`, `ch2` receives `"write: hi"`
- `ReadThenWrite(in, out)` (in="X") → `out` receives `"transform: X"`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/16.channel
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestReadFromBoth
--- PASS: TestReadFromBoth (0.00s)
=== RUN   TestWriteToBoth
--- PASS: TestWriteToBoth (0.00s)
=== RUN   TestReadThenWrite
--- PASS: TestReadThenWrite (0.00s)
PASS
```
