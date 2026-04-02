package sorting

import "sort"

// SortInts sorts a slice of ints in ascending order in-place.
func SortInts(nums []int) {
	sort.Ints(nums)
}

// SortStrings sorts a slice of strings in ascending (lexicographic) order in-place.
func SortStrings(strs []string) {
	sort.Strings(strs)
}

// SortByLength sorts a slice of strings by length (shortest first) in-place.
func SortByLength(strs []string) {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
}

// Person represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// SortByAge sorts a slice of Person by Age in ascending order in-place.
func SortByAge(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

// IsSorted returns true if the given int slice is sorted in ascending order.
func IsSorted(nums []int) bool {
	return sort.IntsAreSorted(nums)
}
