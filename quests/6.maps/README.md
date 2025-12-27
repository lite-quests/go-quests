# Go Maps Quest — Simple Key-Value Cache

Your task is to implement a **basic key-value cache** using Go maps.

This quest focuses on:

- Creating maps with `make`
- Adding and updating entries
- Reading values safely
- Deleting entries
- Iterating over a map
- Returning derived results (not internal state)

---

## Reference

- [https://gobyexample.com/maps](https://gobyexample.com/maps)
- [https://gobyexample.com/range](https://gobyexample.com/range)

---

## Struct

```go
type Cache struct {
	data map[string]int
}
```

---

## Functions to Implement

```go
func NewCache() *Cache
func (c *Cache) Set(key string, value int)
func (c *Cache) Get(key string) (int, bool)
func (c *Cache) Delete(key string)
func (c *Cache) Count() int
func (c *Cache) AllKeys() []string
func (c *Cache) RemoveBelow(limit int)
```

---

## Description

You are building a **simple in-memory cache**.

Each entry maps a `string` key to an `int` value.

The cache starts empty and grows as keys are added.

---

## Function Behavior (Very Explicit)

### 1. `NewCache`

```go
func NewCache() *Cache
```

What to do:

- Create a `Cache`
- Initialize the internal map using `make`
- Return a pointer to the cache

What NOT to do:

- Do not leave the map as `nil`

Example:

```go
c := NewCache()
c.Count() // 0
```

---

### 2. `Set`

```go
func (c *Cache) Set(key string, value int)
```

What to do:

- If `key` does not exist → add it
- If `key` exists → overwrite the value

What to return:

- Nothing

Example:

```go
c.Set("x", 10)
c.Set("x", 20)

c.Get("x") // (20, true)
```

---

### 3. `Get`

```go
func (c *Cache) Get(key string) (int, bool)
```

What to return:

- `(value, true)` if the key exists
- `(0, false)` if the key does NOT exist

Important:

- Do not assume `0` means missing
- Use the `value, ok := map[key]` pattern

Examples:

```go
c.Set("a", 5)

c.Get("a") // (5, true)
c.Get("b") // (0, false)
```

---

### 4. `Delete`

```go
func (c *Cache) Delete(key string)
```

What to do:

- Remove the key from the map if it exists
- Do nothing if the key does not exist

What to return:

- Nothing

Example:

```go
c.Set("a", 1)
c.Delete("a")

c.Get("a") // (0, false)
```

---

### 5. `Count`

```go
func (c *Cache) Count() int
```

What to return:

- The number of key-value pairs in the cache

Example:

```go
c.Set("a", 1)
c.Set("b", 2)

c.Count() // 2
```

---

### 6. `AllKeys`

```go
func (c *Cache) AllKeys() []string
```

What to return:

- A **new slice** containing all keys in the map
- Order does NOT matter

Important:

- Do not return the internal map
- Use `append` to build the slice

Examples:

```go
c.Set("a", 1)
c.Set("b", 2)

keys := c.AllKeys()
// keys contains "a" and "b" (any order)
```

Empty cache:

```go
c := NewCache()
c.AllKeys() // []
```

---

### 7. `RemoveBelow`

```go
func (c *Cache) RemoveBelow(limit int)
```

What to do:

- Iterate over the map
- Delete all entries where `value < limit`

What to return:

- Nothing

Example:

```go
c.Set("a", 5)
c.Set("b", 15)
c.Set("c", 25)

c.RemoveBelow(15)

c.AllKeys() // contains "b", "c"
```

---

## Requirements

1. Use `map[string]int`
2. Use `make` to initialize the map
3. Use `range` at least once
4. Use `delete`
5. Use `len` to count entries
6. Use `append` when building slices
7. Do not expose the internal map

---

## Common Beginner Mistakes (Tests Will Catch These)

- Forgetting to initialize the map
- Returning the internal map directly
- Assuming map iteration order
- Using `0` to detect missing keys
- Forgetting to delete entries inside `RemoveBelow`

---

## Key Ideas (Beginner-Friendly)

- Maps store key-value pairs
- Reading from a map can fail safely
- Maps grow and shrink dynamically
- Iteration order is not guaranteed
- Always return copies, not internals

---

## Run Tests

```bash
go test ./quests/6.maps -v
```
