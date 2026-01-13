package generics

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push adds a value to the end of the list.
func (lst *List[T]) Push(v T) {
	// TODO: implement
}

// Pop removes the last element and returns it.
func (lst *List[T]) Pop() (T, bool) {
	// TODO: implement
	var zero T
	return zero, false
}

// AllElements returns a slice of all elements in the list.
func (lst *List[T]) AllElements() []T {
	// TODO: implement
	return nil
}
