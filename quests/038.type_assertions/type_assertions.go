package type_assertions

import "fmt"

// TODO: Implement the functions below.
// Read README.md for the instructions.

// Describe takes an any value and returns a string describing its type and value.
// Handle the following types: int, float64, string, bool.
// For any other type, return "unknown type".
// Use a type switch.
// Format:
//   int:     "int: 42"
//   float64: "float64: 3.14"
//   string:  "string: hello"
//   bool:    "bool: true"
//   other:   "unknown type"
func Describe(i any) string {
	// TODO: Implement using a type switch
	_ = fmt.Sprintf // hint: use fmt.Sprintf for formatting
	return ""
}

// ExtractInt attempts to extract an int from an any value.
// Returns the int and true if successful, or 0 and false if not.
// Use the comma-ok type assertion idiom.
func ExtractInt(i any) (int, bool) {
	// TODO: Implement using comma-ok assertion
	return 0, false
}

// Stringer is an interface with a String() method.
type Stringer interface {
	String() string
}

// StringifyIfPossible checks if the given value implements the Stringer interface.
// If it does, return the result of calling String() on it.
// If it doesn't, return "not a stringer".
func StringifyIfPossible(i any) string {
	// TODO: Implement using type assertion against the Stringer interface
	return ""
}

// SumInts takes a slice of any values and returns the sum of all int values.
// Non-int values are skipped.
func SumInts(values []any) int {
	// TODO: Implement using a type switch or comma-ok assertion
	return 0
}
