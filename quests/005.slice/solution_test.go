package slice

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the slice Quest 🎉")
	}
	os.Exit(code)
}

func TestProcessScores(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{30, 50, 110, -5, 80}, []int{40, 50, 80}},
		{[]int{-1, -2, 101, 200}, []int{}},
		{[]int{10, 20, 30}, []int{40, 40, 40}},
		{[]int{39, 40, 41}, []int{40, 40, 41}},
		{
			[]int{10, 20, 30, 40, 50, 60, 70},
			[]int{45, 45, 45, 45, 55, 65, 75},
		},
		{
			[]int{95, 96, 97, 98, 99, 100},
			[]int{100, 100, 100, 100, 100, 100},
		},
		{
			[]int{40, 41, 42, 43, 44},
			[]int{40, 41, 42, 43, 44},
		},
		{
			[]int{50, -1, 30, 200, 80, 10},
			[]int{50, 40, 80, 40},
		},
		{
			[]int{20, 20, 20, 20, 20, 20},
			[]int{45, 45, 45, 45, 45, 45},
		},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {

		original := append([]int(nil), tt.input...)

		got := ProcessScores(tt.input)

		if !equalSlices(got, tt.expected) {
			t.Fatalf(
				"ProcessScores(%v) = %v; expected %v",
				tt.input,
				got,
				tt.expected,
			)
		}

		if !equalSlices(tt.input, original) {
			t.Fatalf(
				"ProcessScores mutated input slice: got %v, expected %v",
				tt.input,
				original,
			)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
