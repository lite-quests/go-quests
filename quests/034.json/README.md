# JSON Encoding & Decoding

## Concept

JSON is the lingua franca of web APIs. Go's standard library handles it through the `encoding/json` package. The two core functions are:

- **`json.Marshal(v any) ([]byte, error)`**: Encodes a Go value into JSON bytes.
- **`json.Unmarshal(data []byte, v any) error`**: Decodes JSON bytes into a Go value (pass a pointer).

Go uses **struct tags** to control how fields are serialized:

```go
type Example struct {
    Name    string `json:"name"`           // use "name" as the JSON key
    Age     int    `json:"age,omitempty"`  // omit the field if it's the zero value
    Ignored string `json:"-"`             // never include this field
}
```

For pretty-printed output (useful for debugging or config files), use:

- **`json.MarshalIndent(v any, prefix, indent string) ([]byte, error)`**

## References

- [Go by Example: JSON](https://gobyexample.com/json)
- [pkg.go.dev/encoding/json](https://pkg.go.dev/encoding/json)
- [JSON and Go (official blog)](https://go.dev/blog/json)

## Quest

### Objective

Implement encoding and decoding functions for a `Product` type used in an online store API.

### Requirements

#### Type

A `Product` struct is already defined for you with the following fields and JSON tags:

| Field     | Type      | JSON key     | Notes                        |
|-----------|-----------|--------------|------------------------------|
| `ID`      | `int`     | `"id"`       |                              |
| `Name`    | `string`  | `"name"`     |                              |
| `Price`   | `float64` | `"price"`    |                              |
| `InStock` | `bool`    | `"in_stock"` |                              |
| `Tags`    | `[]string`| `"tags"`     | omitted when nil/empty       |

#### Functions

1. **`MarshalProduct(p Product) ([]byte, error)`**
   - Encode a single `Product` to JSON bytes.

2. **`UnmarshalProduct(data []byte) (Product, error)`**
   - Decode JSON bytes into a `Product`.
   - Return an error if the JSON is invalid.

3. **`MarshalProducts(products []Product) ([]byte, error)`**
   - Encode a slice of `Product` to a JSON array.

4. **`UnmarshalProducts(data []byte) ([]Product, error)`**
   - Decode a JSON array into a slice of `Product`.

5. **`PrettyMarshal(p Product) ([]byte, error)`**
   - Encode a `Product` to indented JSON.
   - Use `""` as prefix and `"\t"` as indent.

### Inputs / Outputs

- Marshal functions take Go values and return `([]byte, error)`.
- Unmarshal functions take `[]byte` and return the Go value plus `error`.

### Examples

```go
p := Product{ID: 1, Name: "Gopher Plushie", Price: 19.99, InStock: true, Tags: []string{"toy"}}

data, _ := MarshalProduct(p)
// data: {"id":1,"name":"Gopher Plushie","price":19.99,"in_stock":true,"tags":["toy"]}

p2, _ := UnmarshalProduct(data)
// p2.Name == "Gopher Plushie"

pretty, _ := PrettyMarshal(p)
// pretty:
// {
// 	"id": 1,
// 	"name": "Gopher Plushie",
// 	...
// }
```

## Tips

- `json.Unmarshal` requires a pointer: `json.Unmarshal(data, &p)`.
- The `omitempty` tag option skips a field when it holds its zero value (`nil` for slices, `0` for ints, `""` for strings, etc.).
- Always check the error returned by both `Marshal` and `Unmarshal`.

## Testing

```bash
go test -v ./quests/034.json
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestMarshalProduct
--- PASS: TestMarshalProduct (0.00s)
=== RUN   TestMarshalProduct_OmitEmptyTags
--- PASS: TestMarshalProduct_OmitEmptyTags (0.00s)
=== RUN   TestUnmarshalProduct
--- PASS: TestUnmarshalProduct (0.00s)
=== RUN   TestUnmarshalProduct_InvalidJSON
--- PASS: TestUnmarshalProduct_InvalidJSON (0.00s)
=== RUN   TestMarshalProducts
--- PASS: TestMarshalProducts (0.00s)
=== RUN   TestUnmarshalProducts
--- PASS: TestUnmarshalProducts (0.00s)
=== RUN   TestPrettyMarshal
--- PASS: TestPrettyMarshal (0.00s)
PASS
```
