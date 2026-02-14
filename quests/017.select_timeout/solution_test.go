package selecttimeout

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
		println(colorGreen, "Success! Completed the select_timeout Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestFunctionOrdered(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	FunctionOrdered()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	got := buf.String()
	want := "from c1\nfrom c4\nfrom c2\nfrom c5\nfrom c3\n"

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}

}
