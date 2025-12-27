package goroutine

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the error Quest 🎉")
	}
	os.Exit(code)
}

func TestSendRequest(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "processed: hello"},
		{"GoLang", "processed: GoLang"},
		{"test123", "processed: test123"},
	}

	for i, tt := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			got := SendRequest(tt.input)
			if got != tt.want {
				t.Fatalf("SendRequest(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
