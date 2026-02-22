# File Handling

## Concept

Go's `os` and `bufio` packages provide powerful tools for reading and writing files. You can open files with `os.Open` for reading and `os.Create` or `os.OpenFile` for writing. The `bufio.Scanner` makes it easy to scan through file content word by word or line by line. When you need to overwrite a file's content entirely, truncating it first with `os.Truncate` or opening with the right flags is essential to avoid leaving stale data behind.

## References

- https://gobyexample.com/reading-files
- https://gobyexample.com/writing-files
- https://pkg.go.dev/os
- https://pkg.go.dev/bufio

## Quest

### Objective

There is a file called `task.txt` sitting in the same folder as your code. It contains a bunch of words. Your job is to implement a single function `ProcessFile()` that:

1. Reads all the words from `task.txt`
2. Capitalizes the **first 5 words** (UPPERCASE) and lowercases the **last 5 words**
3. Writes the modified content back into `task.txt`

How you get there is entirely up to you. You might need a temporary file along the way — if you create one, make sure you clean it up. Or maybe you don't need one at all. Think it through.

### Requirements

#### `ProcessFile` function

```go
func ProcessFile() error
```

- Reads words from `task.txt` (located in the directory `024.files`)
- The **first 5 words** must be converted to **UPPERCASE**
- The **last 5 words** must be converted to **lowercase**
- All words in between remain unchanged
- The result is written back into `task.txt`, words separated by a single space, ending with a newline
- If `task.txt` has fewer than 10 words, return an error: `"task.txt must contain at least 10 words"`
- Any temporary files you create must be cleaned up before the function returns

### Inputs

- No parameters — the file `task.txt` is always in the same directory

### Outputs

- Returns `error` if something goes wrong (file not found, too few words, I/O failure, etc.)
- Returns `nil` on success

### Example

Given `task.txt` before calling `ProcessFile()`:

```
the quick brown fox jumps over the lazy sleeping dog
```

After `ProcessFile()` returns, `task.txt` contains:

```
THE QUICK BROWN FOX JUMPS over the lazy sleeping dog
```

### Constraints

- You may only call `ProcessFile()` — no helper functions need to be exposed
- Internal helpers are fine, keep them unexported
- The first 5 and last 5 words must not overlap — if the file has exactly 10 words, every word gets transformed

**Why This Matters:**

Almost every real Go program touches the filesystem — reading configs, writing logs, transforming data. This quest teaches you to think about the full lifecycle of a file operation: open, read, transform, write back, and clean up after yourself. The cleanup part is just as important as the rest.

## Testing

To run the tests:

```bash
go test -v ./quests/024.files
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestProcessFile
=== RUN   TestProcessFile/less_than_10_words
=== RUN   TestProcessFile/exactly_10_words
=== RUN   TestProcessFile/more_than_10_words_middle_unchanged
--- PASS: TestProcessFile (0.00s)
    --- PASS: TestProcessFile/less_than_10_words (0.00s)
    --- PASS: TestProcessFile/exactly_10_words (0.00s)
    --- PASS: TestProcessFile/more_than_10_words_middle_unchanged (0.00s)
PASS
```
