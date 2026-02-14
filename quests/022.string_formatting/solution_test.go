package stringformatting

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the string formatting Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestStringFormating(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	StringFormating()
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()
	expected := `Product: {Laptop 42 899.99 true}
Details: {name:Laptop quantity:42 price:899.99 available:true}
Go-syntax: main.Product{name:"Laptop", quantity:42, price:899.99, available:true}
Type: main.Product
Available: true
Quantity: 42
Quantity in binary: 101010
Quantity in hex: 2a
Price: $899.99
Formatted: |Laptop    |   42| 899.99|
`

	if got != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, got)
	}

}
