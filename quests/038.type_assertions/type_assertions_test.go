package type_assertions

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the Type Assertions Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		input any
		want  string
	}{
		{42, "int: 42"},
		{3.14, "float64: 3.14"},
		{"hello", "string: hello"},
		{true, "bool: true"},
		{false, "bool: false"},
		{[]int{1, 2}, "unknown type"},
		{nil, "unknown type"},
	}

	for _, tt := range tests {
		got := Describe(tt.input)
		if got != tt.want {
			t.Errorf("Describe(%v): expected %q, got %q", tt.input, tt.want, got)
		}
	}
}

func TestExtractInt(t *testing.T) {
	val, ok := ExtractInt(99)
	if !ok || val != 99 {
		t.Errorf("expected (99, true), got (%d, %v)", val, ok)
	}

	val, ok = ExtractInt("not an int")
	if ok || val != 0 {
		t.Errorf("expected (0, false), got (%d, %v)", val, ok)
	}

	val, ok = ExtractInt(3.14)
	if ok || val != 0 {
		t.Errorf("expected (0, false) for float64, got (%d, %v)", val, ok)
	}
}

// mockStringer implements the Stringer interface for testing.
type mockStringer struct {
	val string
}

func (m mockStringer) String() string {
	return m.val
}

func TestStringifyIfPossible(t *testing.T) {
	got := StringifyIfPossible(mockStringer{"gopher"})
	if got != "gopher" {
		t.Errorf("expected %q, got %q", "gopher", got)
	}

	got = StringifyIfPossible(42)
	if got != "not a stringer" {
		t.Errorf("expected %q, got %q", "not a stringer", got)
	}

	got = StringifyIfPossible("plain string")
	if got != "not a stringer" {
		t.Errorf("expected %q for plain string, got %q", "not a stringer", got)
	}
}

func TestSumInts(t *testing.T) {
	tests := []struct {
		input []any
		want  int
	}{
		{[]any{1, 2, 3}, 6},
		{[]any{1, "skip", 2, 3.14, 4}, 7},
		{[]any{"a", "b", "c"}, 0},
		{[]any{}, 0},
		{[]any{10, 20, 30}, 60},
	}

	for _, tt := range tests {
		got := SumInts(tt.input)
		if got != tt.want {
			t.Errorf("SumInts(%v): expected %d, got %d", tt.input, tt.want, got)
		}
	}
}
