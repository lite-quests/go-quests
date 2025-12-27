# Go Strings & Runes Quest — Unicode-Safe Text Analysis

Your task is to implement **Unicode-safe string processing utilities** in Go.

This quest focuses on:

- Understanding strings as `[]byte`
- Counting bytes vs counting runes
- Iterating over strings using `range`
- Working with rune literals
- Tracking rune positions (byte offsets)
- Avoiding common UTF-8 bugs

---

## Reference

- [https://gobyexample.com/strings-and-runes](https://gobyexample.com/strings-and-runes)
- [https://pkg.go.dev/unicode/utf8](https://pkg.go.dev/unicode/utf8)
- [https://go.dev/blog/strings](https://go.dev/blog/strings)

---

## Struct

```go
type TextStats struct {
	ByteLength int
	RuneCount  int
}
```

---

## Functions to Implement

```go
func AnalyzeText(s string) TextStats
func RuneFrequencies(s string) map[rune]int
func FirstRunePosition(s string, target rune) int
func ExtractRunes(s string) []rune
func HasOnlyASCII(s string) bool
```

---

## Description

You are building a **Unicode-aware text analysis module**.

All functions must work correctly for:

- ASCII text
- UTF-8 text (e.g. Thai, Hindi, emojis)
- Mixed-language strings

Incorrect assumptions about “characters” will cause tests to fail.

---

## Function Behavior (Very Explicit)

---

### 1. `AnalyzeText`

```go
func AnalyzeText(s string) TextStats
```

What to do:

- `ByteLength` → number of **bytes** in the string
- `RuneCount` → number of **Unicode code points**

Requirements:

- Use `len(s)` for bytes
- Use `utf8.RuneCountInString` for runes

Example:

```go
AnalyzeText("hello")
// ByteLength: 5, RuneCount: 5

AnalyzeText("สวัสดี")
// ByteLength: 18, RuneCount: 6
```

---

### 2. `RuneFrequencies`

```go
func RuneFrequencies(s string) map[rune]int
```

What to do:

- Count how many times each rune appears
- Treat the string as Unicode text, not bytes

Requirements:

- Use `range` over the string
- Use `map[rune]int`

Example:

```go
RuneFrequencies("aab")
// map['a':2, 'b':1]

RuneFrequencies("สวัสดีส")
// map['ส':2, 'ว':1, 'ั':1, 'ด':1, 'ี':1]
```

What NOT to do:

- Do not convert to `[]byte`
- Do not assume 1 byte = 1 character

---

### 3. `FirstRunePosition`

```go
func FirstRunePosition(s string, target rune) int
```

What to return:

- The **byte index** where `target` first appears
- `-1` if the rune is not found

Requirements:

- Use `range` to get byte offsets
- Compare using rune literals

Example:

```go
FirstRunePosition("hello", 'l') // 2
FirstRunePosition("สวัสดี", 'ด') // 12
FirstRunePosition("abc", 'x') // -1
```

Important:

- The returned index is a **byte offset**, not rune index

---

### 4. `ExtractRunes`

```go
func ExtractRunes(s string) []rune
```

What to do:

- Return a slice containing all runes in the string
- Order must be preserved

Requirements:

- Use `range`
- Use `append`
- Do NOT manually decode bytes

Example:

```go
ExtractRunes("hi")
// []rune{'h', 'i'}

ExtractRunes("สวัสดี")
// []rune{'ส','ว','ั','ส','ด','ี'}
```

---

### 5. `HasOnlyASCII`

```go
func HasOnlyASCII(s string) bool
```

What to return:

- `true` if **all runes** are ASCII (`<= 127`)
- `false` otherwise

Requirements:

- Iterate over runes, not bytes
- Early return on failure

Example:

```go
HasOnlyASCII("hello")     // true
HasOnlyASCII("hello🙂")  // false
HasOnlyASCII("สวัสดี")   // false
```

---

## Requirements

1. Use `range` over strings
2. Use `utf8.RuneCountInString`
3. Use `map[rune]int`
4. Do not treat strings as `[]byte` unless explicitly required
5. Return new slices/maps (no shared internals)
6. Handle UTF-8 correctly in all cases

---

## Common Beginner Mistakes (Tests Will Catch These)

- Using `len(s)` as character count
- Indexing strings assuming 1 byte per character
- Iterating over `[]byte(s)` instead of runes
- Returning rune index instead of byte offset
- Treating ASCII assumptions as universal

---

## Key Ideas (This Quest Teaches)

- Strings are byte sequences, not characters
- Runes represent Unicode code points
- `range` over strings is Unicode-aware
- Byte offsets and rune indices are different
- UTF-8 safety is explicit, not automatic

---

## Run Tests

```bash
go test ./quests/9.strings -v
```
