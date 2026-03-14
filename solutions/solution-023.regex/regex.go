package regex

import "regexp"

// Read README.md for the instructions

// TODO: Implement IsOnlyNumbers
// IsOnlyNumbers returns true if the string contains only digits (0-9)
func IsOnlyNumbers(s string) bool {
	// 1. Compile pattern: anchored match of one or more digits
	re := regexp.MustCompile(`^[0-9]+$`)
	// 2. Return true if the full string matches
	return re.MatchString(s)
}

// TODO: Implement IsOnlyAlphabets
// IsOnlyAlphabets returns true if the string contains only letters (a-z, A-Z)
func IsOnlyAlphabets(s string) bool {
	// 1. Compile pattern: anchored match of one or more ASCII letters
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	// 2. Return true if the full string matches
	return re.MatchString(s)
}

// TODO: Implement IsValidEmail
// IsValidEmail returns true if the string is a valid email format
func IsValidEmail(s string) bool {
	// 1. Compile pattern: local-part @ domain . TLD (min 2 chars)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	// 2. Return true if the full string matches
	return re.MatchString(s)
}

// TODO: Implement ContainsGoQuest
// ContainsGoQuest returns true if the string contains "go-quest" (case-insensitive)
func ContainsGoQuest(s string) bool {
	// 1. Compile pattern: (?i) flag makes the match case-insensitive
	re := regexp.MustCompile(`(?i)go-quest`)
	// 2. Return true if the string contains "go-quest" in any casing
	return re.MatchString(s)
}

// TODO: Implement IsValidUsername
// IsValidUsername returns true if username is 3-16 characters, alphanumeric plus underscores,
// and must start with a letter
func IsValidUsername(s string) bool {
	// 1. Compile pattern: starts with a letter, followed by 2-15 alphanumeric/underscore chars
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{2,15}$`)
	// 2. Return true if the full string matches
	return re.MatchString(s)
}
