package enum

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
	return ""
}

func NextState(current OrderState, action string) OrderState {
	// TODO: implement
	return current
}
