package enum

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the enums Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestStatesExist(t *testing.T) {
	_ = StateCreated
	_ = StatePaid
	_ = StatePacked
	_ = StateShipped
	_ = StateDelivered
	_ = StateCancelled
	_ = StateReturned
}

func TestOrderStateString(t *testing.T) {
	tests := map[OrderState]string{
		StateCreated:   "created",
		StatePaid:      "paid",
		StatePacked:    "packed",
		StateShipped:   "shipped",
		StateDelivered: "delivered",
		StateCancelled: "cancelled",
		StateReturned:  "returned",
	}

	for state, expected := range tests {
		if state.String() != expected {
			t.Fatalf("expected %q, got %q", expected, state.String())
		}
	}
}

func TestValidTransitions(t *testing.T) {
	tests := []struct {
		from   OrderState
		action string
		to     OrderState
	}{
		{StateCreated, "pay", StatePaid},
		{StatePaid, "pack", StatePacked},
		{StatePacked, "ship", StateShipped},
		{StateShipped, "deliver", StateDelivered},
	}

	for _, tt := range tests {
		got := NextState(tt.from, tt.action)
		if got != tt.to {
			t.Fatalf("from %v with %q: expected %v, got %v", tt.from, tt.action, tt.to, got)
		}
	}
}

func TestTerminalStates(t *testing.T) {
	terminal := []OrderState{
		StateCancelled,
		StateReturned,
	}

	for _, state := range terminal {
		got := NextState(state, "pay")
		if got != state {
			t.Fatalf("terminal state %v should not change", state)
		}
	}
}

func TestUnknownStatePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for unknown state")
		}
	}()

	unknown := OrderState(999)
	NextState(unknown, "pay")
}
