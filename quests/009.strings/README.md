# Go Strings & Runes

## Concept
Strings in Go are immutable sequences of bytes (UTF-8 encoded). A "character" is represented by a `rune` (Unicode code point). Use the `utf8` package and `range` loops to correctly handle multi-byte characters.

## References
- https://gobyexample.com/strings-and-runes
- https://go.dev/blog/strings
- https://pkg.go.dev/unicode/utf8

## Quest

### Objective
Implement Unicode-aware string analysis functions to handle byte vs. rune counting, frequency mapping, position tracking, and ASCII validation.

### Requirements

#### `AnalyzeText`
- Function: `AnalyzeText(s string) TextStats`
- `TextStats` struct: `ByteLength` (int), `RuneCount` (int).
- Populate `ByteLength` with the number of bytes.
- Populate `RuneCount` with the number of Unicode code points.

#### `RuneFrequencies`
- Function: `RuneFrequencies(s string) map[rune]int`
- Return a map where keys are runes and values are their counts.
- Iterate over runes, not bytes.

#### `FirstRunePosition`
- Function: `FirstRunePosition(s string, target rune) int`
- Return the **byte offset** (index) of the first occurrence of `target`.
- Return `-1` if not found.

#### `ExtractRunes`
- Function: `ExtractRunes(s string) []rune`
- Return a slice of all runes in the text, preserving order.

#### `HasOnlyASCII`
- Function: `HasOnlyASCII(s string) bool`
- Return `true` if all runes are ASCII (<= 127).
- Return `false` otherwise (or for mixed content).
- Treat empty string as `true` (all 0 runes are ASCII).

### Inputs
- `s`: Input string (potentially containing multi-byte characters like emojis or non-Latin scripts).
- `target`: Specific rune directly to search for.

### Outputs
- Varies by function (struct, map, index, slice, or boolean).

### Examples
- `AnalyzeText("hi🙂")` → `{ByteLength: 6, RuneCount: 3}`
- `RuneFrequencies("aab")` → `{'a': 2, 'b': 1}`
- `FirstRunePosition("hello", 'l')` → `2`
- `FirstRunePosition("สวัสดี", 'ด')` → `12` (byte offset)
- `ExtractRunes("hi")` → `['h', 'i']`
- `HasOnlyASCII("hello")` → `true`

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/9.strings
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestAnalyzeText
--- PASS: TestAnalyzeText (0.00s)
=== RUN   TestRuneFrequencies
--- PASS: TestRuneFrequencies (0.00s)
=== RUN   TestFirstRunePosition
--- PASS: TestFirstRunePosition (0.00s)
=== RUN   TestExtractRunes
--- PASS: TestExtractRunes (0.00s)
=== RUN   TestHasOnlyASCII
--- PASS: TestHasOnlyASCII (0.00s)
PASS
```
