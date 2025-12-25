package loops

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the Loops Quest 🎉")
	}
	os.Exit(code)
}

func TestSumEvenNumbers(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{6, 12},
		{1, 0},
		{10, 30},
		{0, 0},
		{-1, 0},
	}

	for _, tt := range tests {
		if got := SumEvenNumbers(tt.input); got != tt.expected {
			t.Fatalf(
				"SumEvenNumbers(%d) = %d; expected %d",
				tt.input,
				got,
				tt.expected,
			)
		}
	}
}

func TestKeepOnlyConsonants(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"apple", "banana", "cherry"}, []string{"bnn", "chrry"}},
		{[]string{"aeiou", "xyz"}, []string{"xyz"}},
		{[]string{"hello", "world"}, []string{"hll", "wrld"}},
		{[]string{"AEIOU", "XYZ"}, []string{"XYZ"}},
		{[]string{"", "a", "b"}, []string{"b"}},
	}

	for _, tt := range tests {
		if got := KeepOnlyConsonants(tt.input); !equalSlices(got, tt.expected) {
			t.Fatalf(
				"KeepOnlyConsonants(%v) = %v; expected %v",
				tt.input,
				got,
				tt.expected,
			)
		}
	}
}

func equalSlices(a, b []string) bool {
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
