package sorting

// TODO: Implement the functions below.
// Read README.md for the instructions.

// SortInts sorts a slice of ints in ascending order in-place.
func SortInts(nums []int) {
	// TODO: Implement
}

// SortStrings sorts a slice of strings in ascending (lexicographic) order in-place.
func SortStrings(strs []string) {
	// TODO: Implement
}

// SortByLength sorts a slice of strings by length (shortest first) in-place.
// If two strings have the same length, their relative order does not matter.
func SortByLength(strs []string) {
	// TODO: Implement using sort.Slice
}

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// SortByAge sorts a slice of Person by Age in ascending order in-place.
func SortByAge(people []Person) {
	// TODO: Implement using sort.Slice
}

// IsSorted returns true if the given int slice is sorted in ascending order.
func IsSorted(nums []int) bool {
	// TODO: Implement using sort.IntsAreSorted
	return false
}
