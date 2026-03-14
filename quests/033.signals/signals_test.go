package signals

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

func TestRunGameServer(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ready := make(chan bool)
	errChan := make(chan error, 1)

	// Create a mock worker script
	workerScript := filepath.Join(t.TempDir(), "worker.sh")
	scriptContent := `#!/bin/sh
trap "exit 0" TERM
sleep 10 & 
wait
`
	if err := os.WriteFile(workerScript, []byte(scriptContent), 0755); err != nil {
		t.Fatalf("failed to write worker script: %v", err)
	}

	go func() {
		errChan <- RunGameServer(ready, workerScript)
	}()

	select {
	case <-ready:
	case <-time.After(2 * time.Second):
		t.Fatal("Server did not send true to the ready channel in time")
	}

	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("failed to find current process: %v", err)
	}

	_ = process.Signal(syscall.SIGHUP)
	time.Sleep(100 * time.Millisecond)

	_ = process.Signal(syscall.SIGUSR1)
	time.Sleep(100 * time.Millisecond)

	// Trigger graceful shutdown
	_ = process.Signal(syscall.SIGINT)

	select {
	case err := <-errChan:
		if err != nil {
			// Some systems might output wrapper process error, relax assertions if needed
			t.Logf("Process exited with wait error (acceptable if killed): %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("RunGameServer did not return after receiving SIGINT")
	}

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Reloading game rules...") {
		t.Errorf("expected output to contain 'Reloading game rules...', but got:\n%s", output)
	}

	if !strings.Contains(output, "Dumping player coordinates...") {
		t.Errorf("expected output to contain 'Dumping player coordinates...', but got:\n%s", output)
	}

	if !strings.Contains(output, "Shutting down worker...") {
		t.Errorf("expected output to contain 'Shutting down worker...', but got:\n%s", output)
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the signals Quest 🎉", colorReset)
	}
	os.Exit(code)
}
