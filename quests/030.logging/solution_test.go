package logging_quest

import (
	"bytes"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestProcessSearch(t *testing.T) {
	t.Run("success_valid_query", func(t *testing.T) {
		var buf bytes.Buffer
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		logger := slog.New(slog.NewJSONHandler(&buf, opts))

		ProcessSearch(logger, "req-xyz-123", "running shoes", 42)

		output := buf.String()

		if output == "" {
			t.Fatal("expected log output, got empty string")
		}

		if !strings.Contains(output, `"level":"DEBUG"`) || !strings.Contains(output, `"msg":"processing search request"`) {
			t.Errorf("expected debug log 'processing search request', got %q", output)
		}

		if !strings.Contains(output, `"level":"INFO"`) {
			t.Errorf("expected 'level':'INFO', got %q", output)
		}

		if !strings.Contains(output, `"msg":"search executed"`) {
			t.Errorf("expected 'msg':'search executed', got %q", output)
		}

		if !strings.Contains(output, `"request_id":"req-xyz-123"`) {
			t.Errorf("expected child logger to attach 'request_id':'req-xyz-123', got %q", output)
		}

		if !strings.Contains(output, `"query":"running shoes"`) {
			t.Errorf("expected to find 'query':'running shoes', got %q", output)
		}

		if !strings.Contains(output, `"user":{`) && !strings.Contains(output, `"user": {`) {
			t.Errorf("expected a JSON group for 'user', got %q", output)
		}
		if !strings.Contains(output, `"id":42`) {
			t.Errorf("expected user group to contain 'id':42, got %q", output)
		}
	})

	t.Run("fail_empty_query", func(t *testing.T) {
		var buf bytes.Buffer
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		logger := slog.New(slog.NewJSONHandler(&buf, opts))

		ProcessSearch(logger, "req-empty-456", "", 99)

		output := buf.String()

		if output == "" {
			t.Fatal("expected log output for empty query, got empty string")
		}

		if !strings.Contains(output, `"level":"DEBUG"`) || !strings.Contains(output, `"msg":"processing search request"`) {
			t.Errorf("expected debug log 'processing search request', got %q", output)
		}

		if !strings.Contains(output, `"level":"WARN"`) {
			t.Errorf("expected 'level':'WARN' for empty query, got %q", output)
		}

		if !strings.Contains(output, `"msg":"empty search query"`) {
			t.Errorf("expected 'msg':'empty search query', got %q", output)
		}

		if !strings.Contains(output, `"request_id":"req-empty-456"`) {
			t.Errorf("expected child logger to attach 'request_id':'req-empty-456', got %q", output)
		}

		if strings.Contains(output, `"user":`) {
			t.Errorf("did not expect to see a 'user' group for empty queries, got %q", output)
		}
	})

	t.Run("fail_invalid_user", func(t *testing.T) {
		var buf bytes.Buffer
		opts := &slog.HandlerOptions{Level: slog.LevelDebug}
		logger := slog.New(slog.NewJSONHandler(&buf, opts))

		ProcessSearch(logger, "req-err-789", "shoes", -1)

		output := buf.String()

		if output == "" {
			t.Fatal("expected log output for invalid user, got empty string")
		}

		if !strings.Contains(output, `"level":"DEBUG"`) {
			t.Errorf("expected 'level':'DEBUG' for invalid user query, got %q", output)
		}

		if !strings.Contains(output, `"level":"ERROR"`) {
			t.Errorf("expected 'level':'ERROR' for invalid user, got %q", output)
		}

		if !strings.Contains(output, `"msg":"invalid user ID"`) {
			t.Errorf("expected 'msg':'invalid user ID', got %q", output)
		}

		if !strings.Contains(output, `"request_id":"req-err-789"`) {
			t.Errorf("expected child logger to attach 'request_id':'req-err-789', got %q", output)
		}

		if strings.Contains(output, `"user":`) {
			t.Errorf("did not expect to see a 'user' group for invalid users, got %q", output)
		}
	})
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the logging Quest 🎉", colorReset)
	}
	os.Exit(code)
}
