package solutions

import "unicode/utf8"

type TextStats struct {
	ByteLength int
	RuneCount  int
}

func AnalyzeText(s string) TextStats {
	// TODO: implement
	// Read README.md for the instructions
	var ts TextStats = TextStats{
		ByteLength: len(s),
		RuneCount:  utf8.RuneCountInString(s),
	}
	return ts
}

func RuneFrequencies(s string) map[rune]int {
	// TODO: implement
	// Read README.md for the instructions
	mp := make(map[rune]int)
	for _, r := range s {
		value, ok := mp[r]
		if ok {
			mp[r] = value + 1
		} else {
			mp[r] = 1
		}
	}
	return mp
}

func FirstRunePosition(s string, target rune) int {
	// TODO: implement
	// Read README.md for the instructions
	for i, r := range s {
		if r == target {
			return i
		}
	}
	return -1
}

func ExtractRunes(s string) []rune {

	res := []rune{}
	for _, k := range s {
		res = append(res, k)
	}
	return res
}

func HasOnlyASCII(s string) bool {
	// TODO: implement
	// Read README.md for the instructions
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}
