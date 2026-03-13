package solution023regex

import "regexp"

// IsOnlyNumbers returns true if the string contains only digits (0-9)
func IsOnlyNumbers(s string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(s)
}

// IsOnlyAlphabets returns true if the string contains only letters (a-z, A-Z)
func IsOnlyAlphabets(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	return re.MatchString(s)
}

// IsValidEmail returns true if the string is a valid email format
func IsValidEmail(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}

// ContainsGoQuest returns true if the string contains "go-quest" (case-insensitive)
func ContainsGoQuest(s string) bool {
	re := regexp.MustCompile(`(?i)go-quest`)
	return re.MatchString(s)
}

// IsValidUsername returns true if username is 3-16 characters, alphanumeric plus underscores
// and must start with a letter
func IsValidUsername(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{2,15}$`)
	return re.MatchString(s)
}
