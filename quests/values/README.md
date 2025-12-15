# Go Values Quest

Your task is to implement the function `BuildValues`.

The function should create and return a `Result` struct populated with
different Go value types.

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

Run tests using:

```
go test ./values/solution
```
