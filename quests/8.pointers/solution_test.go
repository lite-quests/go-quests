package pointers

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the functions Quest 🎉")
	}
	os.Exit(code)
}

func TestPointersQuest(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PointersQuest()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()
	want := "15\n0\n9 3\n[1 2 3 100]\n"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}

}
