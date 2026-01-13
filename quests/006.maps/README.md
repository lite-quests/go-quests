# Go Maps

## Concept
Maps are Go's built-in associative data type (hash tables). They map keys to values and provide fast lookups, updates, and deletions. Maps must be initialized using `make` before use.

## References
- https://gobyexample.com/maps
- https://gobyexample.com/range
- https://go.dev/blog/maps

## Quest

### Objective
Implement a simple in-memory key-value cache using a `Cache` struct and Go maps.

### Requirements
- Struct: `type Cache struct { data map[string]int }`
- **Do not expose the internal map directly.**

#### Methods
1. **`NewCache() *Cache`**
   - Initialize the internal map using `make`.
   - Return a pointer to the `Cache`.

2. **`Set(key string, value int)`**
   - Add or update the mapping for `key` to `value`.

3. **`Get(key string) (int, bool)`**
   - Return the value and `true` if exists.
   - Return `0` and `false` if not found.

4. **`Delete(key string)`**
   - Remove the `key` from the map.

5. **`Count() int`**
   - Return the total number of entries.

6. **`AllKeys() []string`**
   - Return a new slice containing all keys (order irrelevant).
   - Return an empty slice if the map is empty.

7. **`RemoveBelow(limit int)`**
   - Iterate through the map and delete all entries where the `value` is strictly less than `limit`.

### Inputs
- Varies by method (see above).

### Outputs
- Varies by method (see above).

### Examples
```go
c := NewCache()
c.Set("a", 10)
c.Set("b", 20)

val, ok := c.Get("a") // 10, true
count := c.Count()    // 2

c.RemoveBelow(15)     // "a" is removed (10 < 15)
keys := c.AllKeys()   // ["b"]
```

## Testing
To run the tests, execute the following command from the root directory:
```bash
go test -v ./quests/6.maps
```

Or from the quest directory:
```bash
go test -v
```
Expected output:
```text
=== RUN   TestCacheOperations
--- PASS: TestCacheOperations (0.00s)
PASS
```
