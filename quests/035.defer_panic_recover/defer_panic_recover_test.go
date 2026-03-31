package defer_panic_recover

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the Defer, Panic & Recover Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestDeferOrder(t *testing.T) {
	got := DeferOrder()
	want := []string{"third", "second", "first"}

	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}

func TestSafeDivide_Normal(t *testing.T) {
	got := SafeDivide(10, 2)
	if got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
}

func TestSafeDivide_PanicsOnZero(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic for division by zero, got none")
		}
		msg, ok := r.(string)
		if !ok {
			t.Errorf("expected panic value to be a string, got %T", r)
		}
		if msg != "division by zero" {
			t.Errorf("expected panic message %q, got %q", "division by zero", msg)
		}
	}()
	SafeDivide(5, 0)
}

func TestRecoverDivide_Normal(t *testing.T) {
	result, panicMsg := RecoverDivide(10, 2)
	if result != 5 {
		t.Errorf("expected result 5, got %d", result)
	}
	if panicMsg != "" {
		t.Errorf("expected empty panicMsg, got %q", panicMsg)
	}
}

func TestRecoverDivide_PanicRecovered(t *testing.T) {
	result, panicMsg := RecoverDivide(10, 0)
	if result != 0 {
		t.Errorf("expected result 0 on panic, got %d", result)
	}
	if panicMsg != "division by zero" {
		t.Errorf("expected panicMsg %q, got %q", "division by zero", panicMsg)
	}
}

func TestCleanup(t *testing.T) {
	got := Cleanup()
	want := []string{"open", "work", "close"}

	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: expected %q, got %q", i, want[i], got[i])
		}
	}
}
