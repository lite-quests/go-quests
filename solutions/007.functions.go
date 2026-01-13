package solutions

import "errors"

// Divide returns a / b or an error if b == 0
func Divide(a, b int) (int, error) {
	if b != 0 {
		c := a / b
		return c, nil
	}

	return 0, errors.New("Divide by 0")
}

// SumAll returns the sum of all provided integers
func SumAll(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// MaxMin returns the max and min of all provided integers
// Returns an error if no numbers are provided
func MaxMin(nums ...int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("Invalid inputs")
	}
	var max int
	var min int

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = nums[i]
			min = nums[i]
			continue
		}
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	return max, min, nil
}

// ConcatAll joins all strings using the provided separator
func ConcatAll(sep string, strs ...string) string {
	var res string = ""
	for i, st := range strs {
		if i == 0 {
			res = st
			continue
		}
		res = res + sep + st
	}
	return res
}
