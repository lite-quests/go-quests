# Context: The Smart File Downloader

## Concept

Think of Go's `context` as a "backpack" that gets passed down the chain from the very first function call (like handling an HTTP request) down to the very last Database query. It carries important request-scoped rules: deadlines, cancellation signals, and metadata (like trace IDs). 

More importantly, Contexts are **hierarchical**. When you create a new context, you derive it from a `parent` context. If the parent gets cancelled, all its children (and their children) get cancelled automatically.

Understanding when to use which type of context is a core Go skill:
1. **`context.Background()`**: The root context. It is never cancelled and has no deadline. Use it at the very top level of an incoming request or main function.
2. **`context.TODO()`**: Identical to `Background()`, but signals to other developers: *"I don't know what context to use here yet, I'll figure it out later."*
3. **`context.WithTimeout(parent, duration)`**: Derives a new context from the given parent that automatically cancels itself after the duration expires. Perfect for network calls.
4. **`context.WithCancel(parent)`**: Derives a new context from the given parent and provides a `cancel()` function. You can manually call `cancel()` to signal workers to stop.
5. **`context.WithValue(parent, key, val)`**: Derives a new context carrying a key-value pair. Used for request-scoped data like trace IDs. (Note: standard Go practice requires using custom types for keys to avoid collisions).

## References
- [Go by Example: Context](https://gobyexample.com/context)
- [pkg.go.dev/context](https://pkg.go.dev/context)

## Quest

### Objective
You are building the core flow for a **Smart File Downloader**. Your CLI tool downloads large assets from a Content Delivery Network (CDN) as fast as possible. You must use all four types of contexts to orchestrate the different steps correctly.

### Requirements

Implement the function:
`func RunSmartDownloader(api *CDNAPI)`

The `api` object (provided by the test suite) has four methods that you must call in this exact order, passing the correct context to each one:

1. **Root Context**: Create an absolute root context using `context.Background()`.
2. **Tracing (`api.LogActivity`)**: The CDN logging system requires a trace ID to track requests.
   - Derive a new context from the root using `context.WithValue`.
   - Use the custom type `TraceKey("trace_id")` for the key and the string `"smart-dl-123"` for the value.
   - Pass this value context to `api.LogActivity(ctx)`.
3. **DNS Lookup (`api.DNSLookup`)**: Next, it does a DNS lookup to find the nearest CDN server. DNS shouldn't take more than `100ms`, so apply a timeout.
   - Derive a new context from the root using `context.WithTimeout` with a duration of `100 * time.Millisecond`.
   - Use `defer` to ensure the timeout's `cancel` function is eventually called (good practice to avoid memory leaks!).
   - Pass this timeout context to `api.DNSLookup(ctx)`.
4. **Concurrent Download (`api.DownloadChunks`)**: Next, it concurrently requests the file chunks from 3 different continent mirrors. Once the file is 100% downloaded, it needs to forcefully close the TCP connections to the other mirrors.
   - Derive a new context from the root using `context.WithCancel`.
   - *Note: Normally, you'd pass this context into a worker goroutine and then defer `cancel()`. For this simulated quest, we just want to verify you know how to extract the cancel func.*
   - Immediately call the `cancel()` function you just received to simulate the file finishing early.
   - Pass the already-cancelled context into `api.DownloadChunks(ctx)`.
5. **Verify Checksum (`api.VerifyChecksum`)**: Finally, it calls a function from a legacy cryptography library that requires a context, but hasn't fully documented its cancellation support yet.
   - Pass a `context.TODO()` into `api.VerifyChecksum(ctx)`.

### Inputs
- `api`: A pointer to a `CDNAPI` instance (simulated by the test suite).

### Outputs
- None (just execution side effects).

## Testing

```bash
go test -v ./quests/1006.context
```
