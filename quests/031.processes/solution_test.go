package processes

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCheckToolVersion(t *testing.T) {
	t.Run("success_version_match", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := CheckToolVersion(ctx, "go", "go") // "go version" output always contains "go"
		if err != nil {
			t.Fatalf("unexpected error checking for 'go': %v", err)
		}
	})

	t.Run("fail_version_mismatch", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := CheckToolVersion(ctx, "go", "version-this-should-not-exist-1234")
		if err == nil {
			t.Fatal("expected error for mismatch version, got nil")
		}

		if !strings.Contains(err.Error(), "version mismatch") {
			t.Errorf("expected error message to contain 'version mismatch', got %q", err.Error())
		}
	})

	t.Run("fail_timeout", func(t *testing.T) {
		// Create a rogue script that sleeps to simulate hanging
		rogueScript := filepath.Join(t.TempDir(), "rogue-tool")
		scriptContent := `#!/bin/sh
sleep 10
`
		if err := os.WriteFile(rogueScript, []byte(scriptContent), 0755); err != nil {
			t.Fatalf("failed to create rogue script: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		start := time.Now()
		err := CheckToolVersion(ctx, rogueScript, "anything")
		duration := time.Since(start)

		if err == nil {
			t.Fatal("expected timeout error, got nil")
		}

		if duration > 2*time.Second {
			t.Fatalf("function hung for %v, expected it to timeout quickly via context", duration)
		}

		if err.Error() != "context deadline exceeded" && !strings.Contains(err.Error(), "killed") && !strings.Contains(err.Error(), "signal") {
			// exec command context typically returns "signal: killed" or "context deadline exceeded" depending on Go version
			t.Logf("Got timeout error as expected: %v", err)
		}
	})
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the processes Quest 🎉", colorReset)
	}
	os.Exit(code)
}
