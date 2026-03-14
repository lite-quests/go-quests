package regex

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the regex Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestIsOnlyNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"12345", true},
		{"abc123", false},
		{"", false},
		{"0000", true},
	}

	for _, test := range tests {
		if IsOnlyNumbers(test.input) != test.expected {
			t.Errorf("IsOnlyNumbers(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}

func TestIsOnlyAlphabets(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"HelloWorld", true},
		{"Hello123", false},
		{"", false},
		{"GoLang", true},
	}

	for _, test := range tests {
		if IsOnlyAlphabets(test.input) != test.expected {
			t.Errorf("IsOnlyAlphabets(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"user@gmail.com", true},
		{"invalid-email", false},
		{"user@domain", false},
		{"xyz@gmail.com", true},
	}

	for _, test := range tests {
		if IsValidEmail(test.input) != test.expected {
			t.Errorf("IsValidEmail(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}

func TestContainsGoQuest(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Welcome to the go-quest!", true},
		{"This is a test string.", false},
		{"GO-QUEST is great!", true},
		{"No mention of the quest here.", false},
	}

	for _, test := range tests {
		if ContainsGoQuest(test.input) != test.expected {
			t.Errorf("ContainsGoQuest(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"user_name", true},
		{"us", false},
		{"this_is_a_very_long_username", false},
		{"validUser123", true},
	}

	for _, test := range tests {
		if IsValidUsername(test.input) != test.expected {
			t.Errorf("IsValidUsername(%q) = %v; want %v", test.input, !test.expected, test.expected)
		}
	}
}
