# Go Worker Pool

## Concept

A worker pool is a concurrency pattern where a fixed number of goroutines (workers) process jobs from a shared channel. Workers process jobs independently and can attach metadata, such as a worker ID, to the output. This pattern is widely used in background processing, batch jobs, and parallel pipelines.

> **Note:** Before attempting this quest, make sure to check out [Go by Example – Worker Pools](https://gobyexample.com/worker-pools). It demonstrates channels, goroutines, and fan-in/fan-out patterns that are critical to solving this quest correctly.

## References

- [https://gobyexample.com/worker-pools](https://gobyexample.com/worker-pools)
- [https://pkg.go.dev/sync](https://pkg.go.dev/sync)
- [https://pkg.go.dev/time#Ticker](https://pkg.go.dev/time#Ticker)

---

## Quest

### Objective

Implement `WorkerPool` to process a list of strings concurrently using **5 workers**, converting each string to **uppercase** and appending the **worker ID** that processed it. Ensure **all jobs are completed before returning**.

---

### Requirements

#### Workers

- Exactly **5 worker goroutines**.
- Each worker must:
  - Read a string from the `inputs` channel.
  - Convert the string to **uppercase**.
  - Append its **worker ID** in the format `-ID` (e.g., `JOB1-2`).
  - Send the result to the `outputs` channel.

#### Channels

- `inputs` channel: For sending jobs to workers.
- `outputs` channel: For collecting processed strings.
- The **length of channels** should be determined by the number of input strings.

#### Behavior

- All jobs must be sent **before closing** the `inputs` channel.
- All results must be received before the function returns.
- Output order **does not need to match input order**.
- No goroutine leaks.

---

### Step-by-Step Guide (for User)

1. Create **inputs and outputs channels** with capacity equal to the number of strings.
2. Start **5 workers** as goroutines, passing each the worker ID and the channels.
3. Send all strings from the input slice into the `inputs` channel.
4. Close the `inputs` channel once all jobs are sent.
5. Receive all results from the `outputs` channel and **append them to a slice**.
6. Return the slice of processed strings.

> Following these steps ensures proper fan-out/fan-in behavior and prevents goroutine leaks.

---

### Output Format

- Slice of strings, each in the format:

```text
JOB1-3
JOB2-1
JOB3-5
...
```

Where:

- `JOBn` is the input string converted to uppercase.
- `-ID` is the worker ID (1–5) that processed the string.

---

### Inputs

- Predefined slice of **100 strings** (e.g., `"job1"`, `"job2"`, … `"job100"`).
- Inputs are **hidden from the user** and provided by the test framework.

---

### Outputs

- Slice of **100 strings**, each in the format `UPPERCASE-WORKERID`.

---

### Examples

If the input slice is:

```text
["job1", "job2", "job3"]
```

A possible output (worker assignment may vary):

```text
["JOB1-2", "JOB2-1", "JOB3-5"]
```

- Worker IDs may differ depending on runtime scheduling.
- Output order does not need to match the input.

---

## Testing

To run the tests, execute the following command from the root directory:

```bash
go test -v ./quests/019.worker_pool
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestWorkerPool
--- PASS: TestWorkerPool (0.03s)
PASS
```
