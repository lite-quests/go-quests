package values

import (
	"fmt"
	"testing"
)

func TestBuildValues(t *testing.T) {
	r := BuildValues()

	if r.Str != "go" {
		t.Errorf("expected Str = go, got %v", r.Str)
	}

	if r.Int != 42 {
		t.Errorf("expected Int = 42, got %v", r.Int)
	}

	if r.Float != 3.14 {
		t.Errorf("expected Float = 3.14, got %v", r.Float)
	}

	if r.Bool != true {
		t.Errorf("expected Bool = true")
	}

	if r.Array != [3]int{1, 2, 3} {
		t.Errorf("array mismatch")
	}

	if len(r.Slice) != 4 || r.Slice[3] != 7 {
		t.Errorf("slice not built correctly")
	}

	if r.Map["apple"] != 2 || r.Map["banana"] != 5 {
		t.Errorf("map values incorrect")
	}

	if r.User.Name != "Alice" || r.User.Age != 20 {
		t.Errorf("user struct incorrect")
	}

	if r.Ptr == nil || *r.Ptr != 10 {
		t.Errorf("pointer incorrect")
	}

	if r.AddFn(3, 4) != 7 {
		t.Errorf("function value not working")
	}

	if v, ok := r.Any.(int); !ok || v != 100 {
		t.Errorf("interface value incorrect")
	}

	if r.ZeroMap != nil {
		t.Errorf("zero map should be nil")
	}
	fmt.Println("Success! Completed the Values Quest 🎉")
}
