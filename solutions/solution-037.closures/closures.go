package closures

// MakeCounter returns a function that increments and returns a counter starting from 0.
func MakeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// MakeAdder returns a function that adds n to any value passed to it.
func MakeAdder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

// MakeMultiplier returns a function that multiplies any value passed to it by factor.
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// MakeAccumulator returns a function that keeps a running total.
func MakeAccumulator() func(int) int {
	total := 0
	return func(n int) int {
		total += n
		return total
	}
}
