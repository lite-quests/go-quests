package strings

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the strings Quest 🎉", colorReset)
	}
	os.Exit(code)
}

/*
========================
AnalyzeText
========================
*/

func TestAnalyzeText(t *testing.T) {
	tests := []struct {
		in   string
		want TextStats
	}{
		{
			in:   "hello",
			want: TextStats{ByteLength: 5, RuneCount: 5},
		},
		{
			in:   "สวัสดี",
			want: TextStats{ByteLength: 18, RuneCount: 6},
		},
		{
			in:   "hi🙂",
			want: TextStats{ByteLength: 6, RuneCount: 3},
		},
		{

			in:   "",
			want: TextStats{ByteLength: 0, RuneCount: 0},
		},
	}

	for _, tt := range tests {
		got := AnalyzeText(tt.in)
		if got != tt.want {
			t.Fatalf("AnalyzeText(%q) = %#v, want %#v", tt.in, got, tt.want)
		}
	}
}

/*
========================
RuneFrequencies
========================
*/

func TestRuneFrequencies(t *testing.T) {
	tests := []struct {
		in   string
		want map[rune]int
	}{
		{

			in:   "aab",
			want: map[rune]int{'a': 2, 'b': 1},
		},
		{

			in:   "สวัสดีส",
			want: map[rune]int{'ส': 3, 'ว': 1, 'ั': 1, 'ด': 1, 'ี': 1},
		},
		{

			in:   "",
			want: map[rune]int{},
		},
	}

	for _, tt := range tests {
		got := RuneFrequencies(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("RuneFrequencies(%q) = %#v, want %#v", tt.in, got, tt.want)
		}
	}
}

/*
========================
FirstRunePosition
========================
*/

func TestFirstRunePosition(t *testing.T) {
	tests := []struct {
		in     string
		target rune
		want   int
	}{
		{

			in:     "hello",
			target: 'l',
			want:   2,
		},
		{

			in:     "สวัสดี",
			target: 'ด',
			want:   12,
		},
		{

			in:     "abc",
			target: 'x',
			want:   -1,
		},
		{

			in:     "",
			target: 'a',
			want:   -1,
		},
	}

	for _, tt := range tests {
		got := FirstRunePosition(tt.in, tt.target)
		if got != tt.want {
			t.Fatalf("FirstRunePosition(%q, %q) = %d, want %d", tt.in, string(tt.target), got, tt.want)
		}
	}
}

/*
========================
ExtractRunes
========================
*/

func TestExtractRunes(t *testing.T) {
	tests := []struct {
		in   string
		want []rune
	}{
		{

			in:   "hi",
			want: []rune{'h', 'i'},
		},
		{

			in:   "สวัสดี",
			want: []rune{'ส', 'ว', 'ั', 'ส', 'ด', 'ี'},
		},
		{

			in:   "a🙂b",
			want: []rune{'a', '🙂', 'b'},
		},
		{

			in:   "",
			want: []rune{},
		},
	}

	for _, tt := range tests {
		got := ExtractRunes(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("ExtractRunes(%q) = %#v, want %#v", tt.in, got, tt.want)
		}
	}
}

/*
========================
HasOnlyASCII
========================
*/

func TestHasOnlyASCII(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{

			in:   "hello",
			want: true,
		},
		{

			in:   "hi🙂",
			want: false,
		},
		{

			in:   "สวัสดี",
			want: false,
		},
		{

			in:   "",
			want: true,
		},
	}

	for _, tt := range tests {
		got := HasOnlyASCII(tt.in)
		if got != tt.want {
			t.Fatalf("HasOnlyASCII(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
