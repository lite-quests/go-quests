# Go Values Quest

Your task is to implement the function `BuildValues`.

The function should create and return a `Result` struct populated with
different Go value types.

## Reference

- https://gobyexample.com/hello-world

### Requirements

1. Set a string value to `"go"`
2. Set an integer to `42`
3. Set a float to `3.14`
4. Set a boolean to `true`
5. Create an array `[3]int` with values `1, 2, 3`
6. Create a slice containing `4, 5, 6, 7`
7. Create a map with:
   - `"apple"` → `2`
   - `"banana"` → `5`
8. Define a struct `User` with:
   - name `"Alice"`
   - age `20`
9. Use a pointer to store the value `10`
10. Assign a function that adds two integers
11. Store an interface value containing the integer `100`
12. Leave a map uninitialized to demonstrate a zero value

Do not print anything. Only return values.

### Go Values – Ultra-Quick Cheat Sheet

**Basic types**

```go
string   → "go"
int      → 42
float64  → 3.14
bool     → true
```

**Composite types**

```go
array  → [3]int{1, 2, 3}        // fixed size
slice  → []int{4, 5, 6}         // dynamic
map    → map[string]int{"a": 1} // key–value
struct → struct{ Name string }{"Mani"}
```

**Pointers**

```go
x := 10
p := &x   // address
*p        // value
```

**Functions as values**

```go
add := func(a, b int) int { return a + b }
```

**Interfaces**

```go
var v interface{}
v = 100
```

**Zero values**

```go
int → 0 | bool → false | string → ""
slice/map/pointer → nil
```

**Key idea:**
Go is statically typed, everything has a clear value and a predictable zero state.

Run tests using:

```bash
go test ./quests/2.values -v
```
