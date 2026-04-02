# Type Assertions & Type Switches

## Concept

In Go, an `interface{}` (or `any` in modern Go) can hold a value of any type. But sometimes you need to get the concrete type back out. Go provides two tools for this:

### Type Assertion
Extracts the concrete value from an interface. There are two forms:

```go
// Form 1: Direct (panics if wrong type)
s := i.(string)

// Form 2: Comma-ok (safe, never panics)
s, ok := i.(string)
if ok {
    fmt.Println("it's a string:", s)
}
```

Always prefer the comma-ok form unless you are 100% certain of the type.

### Type Switch
A cleaner way to handle multiple possible types:

```go
func describe(i any) string {
    switch v := i.(type) {
    case int:
        return fmt.Sprintf("int: %d", v)
    case string:
        return fmt.Sprintf("string: %s", v)
    case bool:
        return fmt.Sprintf("bool: %v", v)
    default:
        return "unknown type"
    }
}
```

You can also assert against interfaces, not just concrete types:

```go
type Stringer interface { String() string }

if s, ok := i.(Stringer); ok {
    fmt.Println(s.String())
}
```

## References

- [Go by Example: Type Assertions](https://gobyexample.com/type-assertions)
- [Go Tour: Type assertions](https://go.dev/tour/methods/15)
- [Go Tour: Type switches](https://go.dev/tour/methods/16)

## Quest

### Objective

Implement four functions that practice type assertions, the comma-ok idiom, type switches, and interface assertions.

### Requirements

#### 1. `Describe(i any) string`
- Use a **type switch** to return a formatted string describing the value.
- Handle: `int`, `float64`, `string`, `bool`.
- For anything else (including `nil`), return `"unknown type"`.
- Format:
  - `int` → `"int: 42"`
  - `float64` → `"float64: 3.14"`
  - `string` → `"string: hello"`
  - `bool` → `"bool: true"`

#### 2. `ExtractInt(i any) (int, bool)`
- Use the **comma-ok** type assertion to safely extract an `int`.
- Return `(value, true)` if the value is an `int`.
- Return `(0, false)` otherwise.

#### 3. `StringifyIfPossible(i any) string`
- Check if `i` implements the `Stringer` interface (already defined: `String() string`).
- If yes, return the result of calling `String()` on it.
- If no, return `"not a stringer"`.

#### 4. `SumInts(values []any) int`
- Iterate over a slice of `any` values.
- Sum only the values that are of type `int`.
- Skip all non-int values.
- Return the total.

### Examples

```go
Describe(42)        // "int: 42"
Describe(3.14)      // "float64: 3.14"
Describe("hello")   // "string: hello"
Describe(true)      // "bool: true"
Describe([]int{})   // "unknown type"

ExtractInt(99)      // (99, true)
ExtractInt("hi")    // (0, false)

// assuming myVal implements Stringer with String() returning "gopher"
StringifyIfPossible(myVal) // "gopher"
StringifyIfPossible(42)    // "not a stringer"

SumInts([]any{1, "skip", 2, 3.14, 4}) // 7
```

## Tips

- In a type switch, the variable `v` in `switch v := i.(type)` takes on the concrete type in each case.
- Use `case nil:` to handle nil explicitly if needed, or let `default` catch it.
- Interface assertions work the same way as concrete type assertions: `val, ok := i.(Stringer)`.

## Testing

```bash
go test -v ./quests/038.type_assertions
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestDescribe
--- PASS: TestDescribe (0.00s)
=== RUN   TestExtractInt
--- PASS: TestExtractInt (0.00s)
=== RUN   TestStringifyIfPossible
--- PASS: TestStringifyIfPossible (0.00s)
=== RUN   TestSumInts
--- PASS: TestSumInts (0.00s)
PASS
```
