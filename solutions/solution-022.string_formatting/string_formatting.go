package string_formatting

import "fmt"

// TODO: Implement StringFormating
// Read README.md for the instructions

type Product struct {
	name      string
	quantity  int
	price     float64
	available bool
}

func StringFormating() {

	// 1. Create a product instance
	p := Product{
		name:      "Laptop",
		quantity:  42,
		price:     899.99,
		available: true,
	}

	// 2. Print the product using default struct formatting
	fmt.Printf("Product: %v\n", p)

	// 3. Print the struct including field names
	fmt.Printf("Details: %+v\n", p)

	// 4. Print the Go-syntax representation of the struct
	fmt.Printf("Go-syntax: %#v\n", p)

	// 5. Print the type of the struct
	fmt.Printf("Type: %T\n", p)

	// 6. Print the availability (boolean)
	fmt.Printf("Available: %t\n", p.available)

	// 7. Print the quantity (decimal)
	fmt.Printf("Quantity: %d\n", p.quantity)

	// 8. Print the quantity in binary
	fmt.Printf("Quantity in binary: %b\n", p.quantity)

	// 9. Print the quantity in hexadecimal
	fmt.Printf("Quantity in hex: %x\n", p.quantity)

	// 10. Print the price formatted with 2 decimal places
	fmt.Printf("Price: $%.2f\n", p.price)

	// 11. Print formatted table-style output with width alignment
	fmt.Printf("Formatted: |%-10s|%5d|%7.2f|\n", p.name, p.quantity, p.price)
}
