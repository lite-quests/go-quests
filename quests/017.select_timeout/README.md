# Go Select & Timeouts

## Concept
The `select` statement lets a goroutine wait on multiple communication operations. It blocks until one of its cases is ready to run. If multiple are ready, it chooses one at random. Timeouts can be implemented using `time.After`.

## References
- https://gobyexample.com/select
- https://gobyexample.com/timeouts
- https://go.dev/tour/concurrency/5

## Quest

### Objective
Implement `FunctionOrdered` to receive messages from 5 concurrently running goroutines in a specific order (`c1`→`c4`→`c2`→`c5`→`c3`) by solely using **timeouts** to control when each channel becomes ready.

### Requirements

#### `FunctionOrdered`
- Create 5 unbuffered channels: `c1` through `c5`.
- Start 5 goroutines, one for each channel.
- Inside each goroutine:
    - Sleep for a specific duration (using `time.Sleep`).
    - Send a specific string to the assigned channel.
- Use a `select` loop (running 5 times) to receive from any ready channel.
- Print the received message immediately.

#### Order & Timing
- **Strict Output Order**: `c1` → `c4` → `c2` → `c5` → `c3`
- **Mechanism**: The order must be achieved by setting appropriate sleep durations in each goroutine.
    - `c1`: Fastest
    - `c4`: 2nd fastest
    - `c2`: 3rd fastest
    - `c5`: 4th fastest
    - `c3`: Slowest

#### Channel Messages
- `c1` sends `"from c1"`
- `c2` sends `"from c2"`
- `c3` sends `"from c3"`
- `c4` sends `"from c4"`
- `c5` sends `"from c5"`

### Inputs
- None.

### Outputs
- Prints 5 lines to stdout in the exact required order.

### Examples
Implementation logic:
```go
go func() {
    time.Sleep(10 * time.Millisecond) // Shortest wait
    c1 <- "from c1"
}()
// ... other goroutines with increasing delays ...
```

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/17.select_timeout
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestFunctionOrdered
from c1
from c4
from c2
from c5
from c3
--- PASS: TestFunctionOrdered (0.11s)
PASS
```
