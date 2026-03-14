package pointers

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
		println(colorGreen, "Success! Completed the pointers Quest 🎉", colorReset)
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
