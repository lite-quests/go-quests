# Go Loops Quest

Your task is to implement two functions using basic looping and string processing in Go.

- `SumEvenNumbers`
- `KeepOnlyConsonants`

Read the requirements carefully and follow idiomatic Go practices.

## Reference

- [https://gobyexample.com/for](https://gobyexample.com/for)
- [https://gobyexample.com/strings](https://gobyexample.com/strings)
- [https://gobyexample.com/range](https://gobyexample.com/range)

---

## Function 1: `SumEvenNumbers`

```go
func SumEvenNumbers(n int) int
```

### Description

Given an integer `n`, compute the sum of **all even numbers from `1` to `n` (inclusive)**.

### Requirements

1. Iterate from `1` to `n`
2. Identify even numbers using modulo (`%`)
3. Accumulate only even values
4. Return the final sum
5. If `n <= 0`, return `0`

### Examples

- `SumEvenNumbers(10)` → `30`
  (2 + 4 + 6 + 8 + 10)
- `SumEvenNumbers(5)` → `6`
  (2 + 4)
- `SumEvenNumbers(-3)` → `0`

Do not print anything. Only return the computed value.

---

## Function 2: `KeepOnlyConsonants`

```go
func KeepOnlyConsonants(strs []string) []string
```

### Description

Given a slice of strings, return a **new slice** where:

- Each string contains **only consonants**
- All vowels (`a, e, i, o, u`) are removed
- Case-insensitive vowel checking (`A` and `a` are both vowels)

### Requirements

1. Iterate over the input slice
2. Process each string character-by-character
3. Remove vowels (`a, e, i, o, u`)
4. Preserve original character case for consonants
5. Maintain the order of strings
6. Return a new slice (do not modify input)

### Examples

- `["hello", "world"]` → `["hll", "wrld"]`
- `["Go", "LANG"]` → `["G", "LNG"]`
- `["aeiou"]` → `[""]`

---

## Go Loops – Ultra-Quick Cheat Sheet

### `for` loop

```go
for i := 0; i < n; i++ {
    // loop body
}
```

### Range over slice

```go
for _, v := range slice {
    // v is the element
}
```

### Range over string (runes)

```go
for _, ch := range str {
    // ch is a rune
}
```

### Modulo check (even number)

```go
if x%2 == 0 {
    // even
}
```

### Rune to lowercase (for vowel check)

```go
unicode.ToLower(ch)
```

---

## Vowels Reference

```go
a, e, i, o, u
```

Everything else (including non-letters) should be treated as consonants unless specified otherwise.

---

### Key Ideas

- Use loops, not recursion
- Prefer clear, readable logic over clever tricks
- Strings in Go are UTF-8; ranging over them yields runes
- Build new values instead of mutating inputs

---

Run tests using:

```bash
go test ./quests/3.loops -v
```

Do not add logging or printing. Your implementation should be pure and deterministic.
