package closures

// TODO: Implement the functions below.
// Read README.md for the instructions.

// MakeCounter returns a function that increments and returns a counter
// starting from 0. Each call to the returned function returns the next value.
// Example: counter := MakeCounter(); counter() == 1; counter() == 2
func MakeCounter() func() int {
	// TODO: Implement
	return nil
}

// MakeAdder returns a function that adds `n` to any value passed to it.
// Example: add5 := MakeAdder(5); add5(3) == 8
func MakeAdder(n int) func(int) int {
	// TODO: Implement
	return nil
}

// MakeMultiplier returns a function that multiplies any value passed to it by `factor`.
// Example: double := MakeMultiplier(2); double(4) == 8
func MakeMultiplier(factor int) func(int) int {
	// TODO: Implement
	return nil
}

// MakeAccumulator returns a function that keeps a running total.
// Each call adds the given value to the total and returns the new total.
// Example: acc := MakeAccumulator(); acc(5) == 5; acc(3) == 8; acc(2) == 10
func MakeAccumulator() func(int) int {
	// TODO: Implement
	return nil
}
