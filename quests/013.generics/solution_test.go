package generics

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the generics Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestElements_Int(t *testing.T) {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	elements := lst.AllElements()
	if len(elements) != 3 {
		t.Fatalf("expected 3 elements, got %d", len(elements))
	}
	if elements[0] != 10 || elements[1] != 20 || elements[2] != 30 {
		t.Fatalf("expected [10 20 30], got %v", elements)
	}

	val, ok := lst.Pop()
	if !ok || val != 30 {
		t.Fatalf("expected Pop() = 30, got %v", val)
	}
	val, ok = lst.Pop()
	if !ok || val != 20 {
		t.Fatalf("expected Pop() = 20, got %v", val)
	}
	val, ok = lst.Pop()
	if !ok || val != 10 {
		t.Fatalf("expected Pop() = 10, got %v", val)
	}

	val, ok = lst.Pop()
	if ok {
		t.Fatalf("expected Pop() on empty list to return !ok, got %v", val)
	}

	elements = lst.AllElements()
	if len(elements) != 0 {
		t.Fatalf("expected empty slice, got %v", elements)
	}
}

func TestElements_String(t *testing.T) {
	lst := List[string]{}
	lst.Push("foo")
	lst.Push("bar")

	elements := lst.AllElements()
	if len(elements) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(elements))
	}
	if elements[0] != "foo" || elements[1] != "bar" {
		t.Fatalf("expected [foo bar], got %v", elements)
	}

	val, ok := lst.Pop()
	if !ok || val != "bar" {
		t.Fatalf("expected Pop() = bar, got %v", val)
	}
	val, ok = lst.Pop()
	if !ok || val != "foo" {
		t.Fatalf("expected Pop() = foo, got %v", val)
	}

	val, ok = lst.Pop()
	if ok {
		t.Fatalf("expected Pop() on empty list to return !ok, got %v", val)
	}

	elements = lst.AllElements()
	if len(elements) != 0 {
		t.Fatalf("expected empty slice, got %v", elements)
	}
}
