package closures

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the Closures Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestMakeCounter(t *testing.T) {
	counter := MakeCounter()

	for i := 1; i <= 5; i++ {
		got := counter()
		if got != i {
			t.Errorf("call %d: expected %d, got %d", i, i, got)
		}
	}
}

func TestMakeCounter_Independent(t *testing.T) {
	c1 := MakeCounter()
	c2 := MakeCounter()

	c1()
	c1()
	c1() // c1 is at 3

	got := c2() // c2 should still be at 1
	if got != 1 {
		t.Errorf("counters should be independent: expected c2() == 1, got %d", got)
	}
}

func TestMakeAdder(t *testing.T) {
	add5 := MakeAdder(5)
	add10 := MakeAdder(10)

	tests := []struct {
		adder func(int) int
		input int
		want  int
	}{
		{add5, 3, 8},
		{add5, 0, 5},
		{add5, -2, 3},
		{add10, 7, 17},
		{add10, -10, 0},
	}

	for _, tt := range tests {
		got := tt.adder(tt.input)
		if got != tt.want {
			t.Errorf("input %d: expected %d, got %d", tt.input, tt.want, got)
		}
	}
}

func TestMakeMultiplier(t *testing.T) {
	double := MakeMultiplier(2)
	triple := MakeMultiplier(3)

	tests := []struct {
		fn    func(int) int
		input int
		want  int
	}{
		{double, 4, 8},
		{double, 0, 0},
		{double, -3, -6},
		{triple, 5, 15},
		{triple, -2, -6},
	}

	for _, tt := range tests {
		got := tt.fn(tt.input)
		if got != tt.want {
			t.Errorf("input %d: expected %d, got %d", tt.input, tt.want, got)
		}
	}
}

func TestMakeAccumulator(t *testing.T) {
	acc := MakeAccumulator()

	steps := []struct {
		add  int
		want int
	}{
		{5, 5},
		{3, 8},
		{2, 10},
		{-4, 6},
		{0, 6},
	}

	for _, s := range steps {
		got := acc(s.add)
		if got != s.want {
			t.Errorf("after adding %d: expected total %d, got %d", s.add, s.want, got)
		}
	}
}

func TestMakeAccumulator_Independent(t *testing.T) {
	acc1 := MakeAccumulator()
	acc2 := MakeAccumulator()

	acc1(100)
	acc1(50) // acc1 total: 150

	got := acc2(10) // acc2 should be independent, total: 10
	if got != 10 {
		t.Errorf("accumulators should be independent: expected acc2(10) == 10, got %d", got)
	}
}
