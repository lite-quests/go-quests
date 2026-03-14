package tickers

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the tickers Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestTicker(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	Ticker()
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()
	want := "hello\nhello\nworld\nhello\nhello\nworld\nhello\nhello\nworld\n"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}

}
