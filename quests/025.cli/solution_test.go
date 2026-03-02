package cli

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the cli Quest 🎉", colorReset)
	}
	os.Exit(code)
}
func captureOutput(f func() error) (string, error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String(), err
}

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	file := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	return file
}

func TestRunCLI_Count(t *testing.T) {
	content := "line1\nline2\nline3\nline4\nline5"
	file := createTempFile(t, content)
	args := []string{"app", "count", file}

	output, err := captureOutput(func() error {
		return RunCLI(args)
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "lines: 5"

	if !strings.Contains(output, expected) {
		t.Errorf("expected output to contain %q, got %q", expected, output)
	}
}

func TestRunCLI_Search(t *testing.T) {
	content := "Hello World\n Go is great\nhello there\nLearning Go"
	file := createTempFile(t, content)

	t.Run("search case-sensitive", func(t *testing.T) {
		args := []string{"app", "search", "Hello", file}
		output, err := captureOutput(func() error {
			return RunCLI(args)
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.Contains(output, "Hello World") {
			t.Errorf("expected output to contain 'Hello World, got %q", output)
		}

		if strings.Contains(output, "hello there") {
			t.Errorf("expected output NOT to contain 'hello there' (case sensitive), got %q", output)
		}
	})

	t.Run("search case-insensitive", func(t *testing.T) {
		args := []string{"app", "search", "--case-insensitive", "hello", file}
		output, err := captureOutput(func() error {
			return RunCLI(args)
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.Contains(output, "Hello World") {
			t.Errorf("expected output to contain 'Hello World', got %q", output)
		}

		if !strings.Contains(output, "hello there") {
			t.Errorf("expected output to contain 'hello there', got %q", output)
		}
	})
}

func TestRunCLI_Errors(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		err := RunCLI([]string{"app"})
		if err == nil {
			t.Error("expected error for no args, got nil")
		}
	})

	t.Run("unknown subcommand", func(t *testing.T) {
		err := RunCLI([]string{"app", "unknown"})
		if err == nil {
			t.Error("expected error for unknown subcommand, got nil")
		}
	})

	t.Run("missing file arg for count", func(t *testing.T) {
		err := RunCLI([]string{"app", "count"})
		if err == nil {
			t.Error("expected error for missing file arg, got nil")
		}
	})

	t.Run("missing file arg for serach", func(t *testing.T) {
		err := RunCLI([]string{"app", "search", "term"})
		if err == nil {
			t.Error("expected error for missing file arg, got nil")
		}
	})
}
