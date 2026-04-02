package defer_panic_recover

// DeferOrder appends strings to a slice using defer to demonstrate LIFO order.
func DeferOrder() (result []string) {
	result = []string{}
	defer func() { result = append(result, "first") }()
	defer func() { result = append(result, "second") }()
	defer func() { result = append(result, "third") }()
	return
}

// SafeDivide divides a by b, panicking if b is 0.
func SafeDivide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// RecoverDivide calls SafeDivide and recovers from any panic.
func RecoverDivide(a, b int) (result int, panicMsg string) {
	defer func() {
		if r := recover(); r != nil {
			panicMsg = r.(string)
			result = 0
		}
	}()
	result = SafeDivide(a, b)
	return result, ""
}

// Cleanup demonstrates defer for resource cleanup.
func Cleanup() (log []string) {
	log = []string{}
	log = append(log, "open")
	defer func() { log = append(log, "close") }()
	log = append(log, "work")
	return
}
