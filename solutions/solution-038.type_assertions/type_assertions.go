package type_assertions

import "fmt"

// Describe takes an any value and returns a string describing its type and value.
func Describe(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case float64:
		return fmt.Sprintf("float64: %g", v)
	case string:
		return fmt.Sprintf("string: %s", v)
	case bool:
		return fmt.Sprintf("bool: %v", v)
	default:
		return "unknown type"
	}
}

// ExtractInt attempts to extract an int from an any value using comma-ok.
func ExtractInt(i any) (int, bool) {
	v, ok := i.(int)
	return v, ok
}

// Stringer is an interface with a String() method.
type Stringer interface {
	String() string
}

// StringifyIfPossible checks if the value implements Stringer and calls String().
func StringifyIfPossible(i any) string {
	if s, ok := i.(Stringer); ok {
		return s.String()
	}
	return "not a stringer"
}

// SumInts sums all int values in a slice of any, skipping non-int values.
func SumInts(values []any) int {
	total := 0
	for _, v := range values {
		if n, ok := v.(int); ok {
			total += n
		}
	}
	return total
}
