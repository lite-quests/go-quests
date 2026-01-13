package tickers

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the tickers Quest 🎉")
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
