package hello_world

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
		println(colorGreen, "Success! Completed the Hello Go Quest 🎉", colorReset)

	}
	os.Exit(code)
}

func TestHelloWorld(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	HelloGo()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()
	want := "Yo! Hello Go"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
