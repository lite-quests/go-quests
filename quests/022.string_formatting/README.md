# String Formatting

## Concept

String formatting in Go is handled by the `fmt` package, which provides printf-style formatting for displaying values. Format verbs (like `%d`, `%s`, `%f`) control how different data types are printed. Beyond basic printing, you can control width, padding, precision, and alignment to create well-formatted output. Understanding these verbs is essential for debugging, logging, and creating user-friendly displays.

## References

- https://gobyexample.com/string-formatting
- https://pkg.go.dev/fmt
- https://go.dev/blog/strings

## Quest

### Objective

Create a program that formats and displays information about a product using various format verbs. You'll need to explore the references to find the right format verbs for each output requirement.

### Requirements

#### Declare a `Product` struct

```go
type Product struct {
    name      string
    quantity  int
    price     float64
    available bool
}
```

#### Create a product instance

```go
p := Product{
    name:      "Laptop",
    quantity:  42,
    price:     899.99,
    available: true,
}
```

#### Print 10 formatted lines

Your program must produce **exactly** these 10 lines of output:

```text
Product: {Laptop 42 899.99 true}
Details: {name:Laptop quantity:42 price:899.99 available:true}
Go-syntax: main.Product{name:"Laptop", quantity:42, price:899.99, available:true}
Type: main.Product
Available: true
Quantity: 42
Quantity in binary: 101010
Quantity in hex: 2a
Price: $899.99
Formatted: |Laptop    |   42| 899.99|
```

### Inputs

- No external inputs required
- Use the declared Product struct and instance

### Outputs

Your program must print exactly 10 lines matching the format shown above. Study the references to discover which format verbs produce each output style.

**Why String Formatting Matters:**

- **Debugging**: Different struct formats show varying levels of detail during development
- **User interfaces**: Control decimal places and alignment for readable displays
- **Logging**: Include type information and convert numbers to different bases for analysis
- **Professional output**: Width and alignment create clean, table-like displays

Learning to choose the right format verb for each situation is a core Go skill.

## Testing

To run the tests:

```bash
go test -v ./quests/022.string_formatting
```

Or from the quest directory:

```bash
go test -v
```

Expected output:

```text
=== RUN   TestStringFormating
--- PASS: TestStringFormating (0.00s)
PASS
```
