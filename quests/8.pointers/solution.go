package pointers

import "fmt"

func addTenValue(n int) {
	n += 10
}

func addTenPointer(n *int) {
	*n += 10
}

// -------- Task 2 --------

func resetValue(n int) {
	n = 0
}

func resetPointer(n *int) {
	*n = 0
}

// -------- Task 3 --------

func swapValue(a, b int) {
	a, b = b, a
}

func swapPointer(a, b *int) {
	*a, *b = *b, *a
}

// -------- Task 4 --------

func appendValue(s []int) {
	s = append(s, 100)
}

func appendPointer(s *[]int) {
	*s = append(*s, 100)
}

func PointersQuest() {
	// Task 1:
	// Make x become 15
	x := 5
	// CALL THE CORRECT FUNCTION HERE
	fmt.Println(x) // expect: 15

	// Task 2:
	// Make y become 0
	y := 42
	// CALL THE CORRECT FUNCTION HERE
	fmt.Println(y) // expect: 0

	// Task 3:
	// Swap a and b so output is: 9 3
	a, b := 3, 9
	// CALL THE CORRECT FUNCTION HERE
	fmt.Println(a, b) // expect: 9 3

	// Task 4:
	// Make slice contain [1 2 3 100]
	s := []int{1, 2, 3}
	// CALL THE CORRECT FUNCTION HERE
	fmt.Println(s) // expect: [1 2 3 100]
}
