package solutions

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
	e := &element[T]{val: v}
	if lst.head == nil {
		lst.head = e
		lst.tail = e
		return
	}

	lst.tail.next = e
	lst.tail = e
}

// Pop removes the last element and returns it.
func (lst *List[T]) Pop() (T, bool) {
	// TODO: implement
	var zero T
	if lst.head == nil {
		return zero, false
	}
	e := lst.tail.val
	if lst.head == lst.tail {
		lst.head = nil
		lst.tail = nil
	} else {
		curr := lst.head
		for curr.next != lst.tail {
			curr = curr.next
		}
		curr.next = nil
		lst.tail = curr
	}
	return e, true
}

// AllElements returns a slice of all elements in the list.
func (lst *List[T]) AllElements() []T {
	// TODO: implement
	var res []T

	for curr := lst.head; curr != nil; curr = curr.next {
		res = append(res, curr.val)
	}
	return res
}
