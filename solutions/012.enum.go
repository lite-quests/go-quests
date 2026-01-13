package solutions

import "fmt"

type OrderState int

const (
	StateCreated OrderState = iota
	StatePaid
	StatePacked
	StateShipped
	StateDelivered
	StateCancelled
	StateReturned
)

func (s OrderState) String() string {
	// TODO: implement
	var mp = map[OrderState]string{
		StateCreated:   "created",
		StatePaid:      "paid",
		StatePacked:    "packed",
		StateShipped:   "shipped",
		StateDelivered: "delivered",
		StateCancelled: "cancelled",
		StateReturned:  "returned",
	}

	return mp[s]
}

func NextState(current OrderState, action string) OrderState {
	// TODO: implement
	switch current {
	case StateCreated:
		switch action {
		case "pay":
			return StatePaid
		case "cancel":
			return StateCancelled
		}
		return current

	case StatePaid:
		switch action {
		case "pack":
			return StatePacked
		case "cancel":
			return StateCancelled
		}
		return current

	case StatePacked:
		if action == "ship" {
			return StateShipped
		}
		return current

	case StateShipped:
		switch action {
		case "deliver":
			return StateDelivered
		case "return":
			return StateReturned
		}
		return current

	case StateDelivered:
		if action == "return" {
			return StateReturned
		}
		return current

	case StateCancelled, StateReturned:
		return current
	default:
		panic(fmt.Errorf("unknown state: %s", current))
	}
}
