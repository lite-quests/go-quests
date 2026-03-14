package processes

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// TODO: Implement CheckToolVersion
// Read README.md for the instructions

func CheckToolVersion(ctx context.Context, tool string, versionString string) error {
	// 1. Use exec.CommandContext to prepare the execution of the tool with the argument "version"
	cmd := exec.CommandContext(ctx, tool, "version")

	// Put the process in its own group so we can kill it and all its children on timeout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	// Kill the entire process group when the context is cancelled, preventing pipe hangs
	cmd.Cancel = func() error {
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}

	// 2. Call .Output() to execute the command and capture its standard output
	output, err := cmd.Output()

	// 3. If .Output() returns an error (timeout, non-zero exit, or missing executable), return it as-is
	if err != nil {
		return err
	}

	// 4. Convert the returned output byte slice into a string
	// 5. Strip any whitespace/newlines from the ends to keep the error message clean
	outputString := strings.TrimSpace(string(output))

	// 5. If the output does not contain the expected versionString, return a version mismatch error
	if !strings.Contains(outputString, versionString) {
		return fmt.Errorf("version mismatch: expected %s, got %s", versionString, outputString)
	}

	// 6. Output contains the versionString, return nil
	return nil
}
