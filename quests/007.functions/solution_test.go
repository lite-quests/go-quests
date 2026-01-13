package functions

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the functions Quest 🎉")
	}
	os.Exit(code)
}

func TestFunctions(t *testing.T) {
	// --- Test Divide ---
	divTests := []struct {
		a, b        int
		want        int
		expectError bool
	}{
		{10, 2, 5, false},
		{20, 4, 5, false},
		{10, 0, 0, true},
	}

	for _, tt := range divTests {
		got, err := Divide(tt.a, tt.b)
		if (err != nil) != tt.expectError || got != tt.want {
			t.Fatalf("Divide(%d, %d) = (%v, %v); expected (%v, error=%v)", tt.a, tt.b, got, err, tt.want, tt.expectError)
		}
	}

	// --- Test SumAll ---
	sumTests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{[]int{5, -2, 7}, 10},
	}

	for _, tt := range sumTests {
		got := SumAll(tt.input...)
		if got != tt.want {
			t.Fatalf("SumAll(%v) = %v; expected %v", tt.input, got, tt.want)
		}
	}

	// --- Test MaxMin ---
	maxMinTests := []struct {
		input       []int
		wantMax     int
		wantMin     int
		expectError bool
	}{
		{[]int{1, 5, 3}, 5, 1, false},
		{[]int{10}, 10, 10, false},
		{[]int{}, 0, 0, true},
	}

	for _, tt := range maxMinTests {
		gotMax, gotMin, err := MaxMin(tt.input...)
		if (err != nil) != tt.expectError || gotMax != tt.wantMax || gotMin != tt.wantMin {
			t.Fatalf("MaxMin(%v) = (%v, %v, %v); expected (%v, %v, error=%v)", tt.input, gotMax, gotMin, err, tt.wantMax, tt.wantMin, tt.expectError)
		}
	}

	// --- Test ConcatAll ---
	concatTests := []struct {
		sep   string
		input []string
		want  string
	}{
		{"-", []string{"a", "b", "c"}, "a-b-c"},
		{",", []string{"x", "y"}, "x,y"},
		{":", []string{}, ""},
	}

	for _, tt := range concatTests {
		got := ConcatAll(tt.sep, tt.input...)
		if got != tt.want {
			t.Fatalf("ConcatAll(%q, %v) = %q; expected %q", tt.sep, tt.input, got, tt.want)
		}
	}
}
