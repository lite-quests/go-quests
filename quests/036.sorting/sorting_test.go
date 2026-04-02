package sorting

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the Sorting Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestSortInts(t *testing.T) {
	nums := []int{5, 2, 8, 1, 9, 3}
	SortInts(nums)
	want := []int{1, 2, 3, 5, 8, 9}
	for i, v := range want {
		if nums[i] != v {
			t.Errorf("index %d: expected %d, got %d", i, v, nums[i])
		}
	}
}

func TestSortInts_AlreadySorted(t *testing.T) {
	nums := []int{1, 2, 3}
	SortInts(nums)
	want := []int{1, 2, 3}
	for i, v := range want {
		if nums[i] != v {
			t.Errorf("index %d: expected %d, got %d", i, v, nums[i])
		}
	}
}

func TestSortStrings(t *testing.T) {
	strs := []string{"banana", "apple", "cherry", "avocado"}
	SortStrings(strs)
	want := []string{"apple", "avocado", "banana", "cherry"}
	for i, v := range want {
		if strs[i] != v {
			t.Errorf("index %d: expected %q, got %q", i, v, strs[i])
		}
	}
}

func TestSortByLength(t *testing.T) {
	strs := []string{"banana", "fig", "kiwi", "go", "pineapple"}
	SortByLength(strs)

	// Verify lengths are non-decreasing
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(strs[i-1]) {
			t.Errorf("not sorted by length at index %d: %q (len %d) comes after %q (len %d)",
				i, strs[i], len(strs[i]), strs[i-1], len(strs[i-1]))
		}
	}
}

func TestSortByAge(t *testing.T) {
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Diana", Age: 28},
	}
	SortByAge(people)
	wantAges := []int{25, 28, 30, 35}
	for i, age := range wantAges {
		if people[i].Age != age {
			t.Errorf("index %d: expected age %d, got %d", i, age, people[i].Age)
		}
	}
}

func TestIsSorted_True(t *testing.T) {
	if !IsSorted([]int{1, 2, 3, 4, 5}) {
		t.Error("expected true for sorted slice")
	}
}

func TestIsSorted_False(t *testing.T) {
	if IsSorted([]int{3, 1, 2}) {
		t.Error("expected false for unsorted slice")
	}
}

func TestIsSorted_SingleElement(t *testing.T) {
	if !IsSorted([]int{42}) {
		t.Error("expected true for single-element slice")
	}
}
