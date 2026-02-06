package exit_quest

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

const helperEnvVar = "GO_WANT_EXIT_PROCESS"

func TestExitWithStatus(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperProcess")

	cmd.Env = append(os.Environ(), helperEnvVar+"=1")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err == nil {
		t.Fatal("expected process to exit, but it did not")
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("expected ExitError, got %T", err)
	}

	exitCode := exitErr.ExitCode()

	if exitCode != 3 {
		t.Fatalf("expected exit code 3, got %d", exitCode)
	}

	output := stdout.String() + stderr.String()

	if output != "" {
		t.Fatalf("expected no output, but got %q", output)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv(helperEnvVar) != "1" {
		return
	}

	ExitWithStatus()

	os.Exit(0)
}
