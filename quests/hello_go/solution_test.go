package hello_world

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

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
	fmt.Println("Success! Completed the Hello Go Quest 🎉")
}
