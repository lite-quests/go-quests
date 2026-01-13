package channel

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		println("Success! Completed the channel Quest 🎉")
	}
	os.Exit(code)
}

func TestReadFromBoth(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "message from ch1"
	}()

	go func() {
		ch2 <- "message from ch2"
	}()

	msg := ReadFromBoth(ch1, ch2)
	if msg != "read: message from ch1 & message from ch2" {
		t.Fatalf("ReadFromBoth() = %q; want \"read: message from ch1 & message from ch2\"", msg)
	}
}

func TestWriteToBoth(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	WriteToBoth(ch1, ch2, "hello")
	msg1 := <-ch1
	msg2 := <-ch2
	if msg1 != "write: hello" {
		t.Fatalf("WriteToBoth() failed for to write message to ch1 got %q; want \"write: hello\"", msg1)
	}
	if msg2 != "write: hello" {
		t.Fatalf("WriteToBoth() failed for to write message to ch2 got %q; want \"write: hello\"", msg2)
	}
}

func TestReadThenWrite(t *testing.T) {
	input := make(chan string)
	output := make(chan string)
	go func() {
		input <- "msg"
	}()

	go ReadThenWrite(input, output)
	msg := <-output
	if msg != "transform: msg" {
		t.Fatalf("ReadThenWrite() = %q; want \"transform: msg\"", msg)
	}
}
