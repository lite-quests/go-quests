package conditions

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the conditions Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestClassifyRequest(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		hasID    bool
		balance  float64
		isVIP    bool
		expected string
	}{
		{
			name:     "invalid_negative_age",
			age:      -1,
			hasID:    true,
			balance:  1000,
			isVIP:    false,
			expected: "INVALID",
		},
		{
			name:     "invalid_negative_balance",
			age:      25,
			hasID:    true,
			balance:  -10,
			isVIP:    false,
			expected: "INVALID",
		},
		{
			name:     "rejected_underage",
			age:      17,
			hasID:    true,
			balance:  5000,
			isVIP:    true,
			expected: "REJECTED",
		},
		{
			name:     "rejected_no_id",
			age:      30,
			hasID:    false,
			balance:  5000,
			isVIP:    true,
			expected: "REJECTED",
		},
		{
			name:     "vip_access",
			age:      35,
			hasID:    true,
			balance:  15000,
			isVIP:    true,
			expected: "VIP_ACCESS",
		},
		{
			name:     "standard_access",
			age:      40,
			hasID:    true,
			balance:  3000,
			isVIP:    false,
			expected: "STANDARD_ACCESS",
		},
		{
			name:     "limited_access_low_balance",
			age:      22,
			hasID:    true,
			balance:  200,
			isVIP:    false,
			expected: "LIMITED_ACCESS",
		},
		{
			name:     "vip_flag_but_insufficient_balance",
			age:      28,
			hasID:    true,
			balance:  5000,
			isVIP:    true,
			expected: "STANDARD_ACCESS",
		},
		{
			name:     "boundary_exact_1000",
			age:      18,
			hasID:    true,
			balance:  1000,
			isVIP:    false,
			expected: "STANDARD_ACCESS",
		},
		{
			name:     "boundary_exact_10000_vip",
			age:      21,
			hasID:    true,
			balance:  10000,
			isVIP:    true,
			expected: "VIP_ACCESS",
		},
	}

	for _, tt := range tests {
		if got := ClassifyRequest(tt.age, tt.hasID, tt.balance, tt.isVIP); got != tt.expected {
			t.Fatalf(
				"%s: ClassifyRequest(%d, %v, %.2f, %v) = %s; expected %s",
				tt.name,
				tt.age,
				tt.hasID,
				tt.balance,
				tt.isVIP,
				got,
				tt.expected,
			)
		}
	}
}

func TestEvaluateGrade(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{55, "F"},
		{100, "A"},
		{0, "F"},
		{89, "B"},
		{79, "C"},
		{69, "D"},
		{-1, "INVALID"},
		{101, "INVALID"},
	}

	for _, tt := range tests {
		if got := EvaluateGrade(tt.score); got != tt.expected {
			t.Fatalf(
				"EvaluateGrade(%d) = %s; expected %s",
				tt.score,
				got,
				tt.expected,
			)
		}
	}
}
