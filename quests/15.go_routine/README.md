# Go Goroutines & Channels Quest — Simple Server Call

In Go, **concurrency is not about threads**.
It is about **doing work independently and communicating safely**.

Go gives us two tools for this:

- **Goroutines** → run work in the background
- **Channels** → pass data between running code

This quest helps you understand how these two work **together**.

---

## What Is a Goroutine?

A goroutine is a function that runs **independently** of the caller.

Normally, when you call a function, Go waits for it to finish:

```go
f("hello")
```

If you add the `go` keyword:

```go
go f("hello")
```

Go starts running `f` **in the background** and immediately moves on.

The important thing to understand:

> When a goroutine runs, the caller does **not wait** for it.

So if you need a result back, you must use something else.

That “something else” is a **channel**.

---

## What Is a Channel?

A channel is a typed pipe that lets goroutines **send values to each other**.

Think of it like this:

- One goroutine puts data **into** the channel
- Another goroutine takes data **out of** the channel

Example:

```go
ch := make(chan string)

go func() {
    ch <- "hello"
}()

msg := <-ch
```

Here’s the key idea:

> Sending blocks until someone receives
> Receiving blocks until someone sends

Because of this, channels also act as **synchronization tools**.

---

## What This Quest Is About

You are going to build a **very small server-like flow**.

The flow looks like this:

1. A client sends a string request
2. A server processes it in a goroutine
3. The server sends the result back through a channel
4. The client waits and receives the result

This models how real systems work internally:

- background jobs
- services
- async processing

---

## The Goal

You will write code where:

- A string is passed to a function
- That function starts a goroutine
- The goroutine processes the string
- The result is returned using a channel

No sleeps.
No wait groups.
Just goroutines and channels.

---

## Functions You Must Implement

### 1. Client Function

```go
func SendRequest(input string) string
```

This function represents the **caller**.

What it should do:

- Create a channel
- Start the server using a goroutine
- Wait for the response from the channel
- Return the response

This function should **block naturally** while waiting.

---

### 2. Server Function

```go
func Server(request string, responseCh chan string)
```

This function represents the **server**.

What it should do:

- Run inside a goroutine
- Process the request
- Send exactly one string into the channel

The server should not receive from the channel and should not close it.

---

## Processing Rule

The server must produce the response in this exact format:

```text
processed: <input>
```

Example:

```text
input:  "hello"
output: "processed: hello"
```

No extra spaces, no formatting changes.

---

## How the Code Should Feel

When written correctly:

- The client does not know _how_ the server runs
- The server does not know _who_ is waiting
- The channel connects them safely

This is idiomatic Go.

---

## Example Usage

```go
result := SendRequest("ping")
fmt.Println(result)
```

Expected output:

```text
processed: ping
```

---

## Important Rules

- Use **one goroutine**
- Use **one unbuffered channel**
- Do not use `time.Sleep`
- Do not use `sync.WaitGroup`
- Let the channel handle waiting
- Do not use global variables

If the program works, it means your understanding is correct.

---

## Why This Works

When the client tries to receive from the channel:

```go
result := <-responseCh
```

Go will **pause execution** until the server sends a value.

This is why we don’t need sleeps or locks.

The channel is doing the coordination for us.

---

## References (Read These Slowly)

These explain the same ideas with examples:

- [https://gobyexample.com/goroutines](https://gobyexample.com/goroutines)
- [https://gobyexample.com/channels](https://gobyexample.com/channels)
- [https://go.dev/doc/effective_go#concurrency](https://go.dev/doc/effective_go#concurrency)

---

## Run the Tests

Once your solution is ready, run:

```bash
go test ./quests/15.go_routine -v
```
