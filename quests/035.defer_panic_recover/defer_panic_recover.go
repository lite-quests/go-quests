package defer_panic_recover

// TODO: Implement the functions below.
// Read README.md for the instructions.

// DeferOrder appends strings to a slice using defer to demonstrate LIFO order.
// Defer three calls that append "first", "second", "third" to the slice,
// then return the slice after all defers have run.
func DeferOrder() []string {
	result := []string{}
	// TODO: Implement using defer
	return result
}

// SafeDivide divides a by b and returns the result.
// If b is 0, it panics with the message "division by zero".
func SafeDivide(a, b int) int {
	// TODO: Implement (panic when b == 0)
	return 0
}

// RecoverDivide calls SafeDivide(a, b) and recovers from any panic.
// If a panic occurs, return 0 and the recovered value as a string.
// If no panic, return the result and an empty string.
func RecoverDivide(a, b int) (result int, panicMsg string) {
	// TODO: Implement using defer + recover
	return 0, ""
}

// Cleanup simulates opening and closing a resource.
// It appends "open" to log, then defers appending "close" to log.
// Then it appends "work" to log.
// Returns the log slice after all defers run.
func Cleanup() []string {
	log := []string{}
	// TODO: Implement
	return log
}
