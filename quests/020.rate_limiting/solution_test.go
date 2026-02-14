package ratelimiting

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the rate-limiting Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestPingPong(t *testing.T) {
	tests := []struct {
		name         string
		requests     []string
		minDuration  time.Duration
		maxDuration  time.Duration
		expectedResp []string
	}{
		{
			name:         "test_only_burst",
			requests:     []string{"Ping", "Pong", "Ping", "Pong"},
			minDuration:  0 * time.Millisecond,
			maxDuration:  50 * time.Millisecond,
			expectedResp: []string{"Pong", "Ping", "Pong", "Ping"},
		},
		{
			name:         "test_burst_then_rate_limit",
			requests:     []string{"Ping", "Pong", "Ping", "Pong", "Ping", "Pong", "Ping"},
			minDuration:  250 * time.Millisecond,
			maxDuration:  400 * time.Millisecond,
			expectedResp: []string{"Pong", "Ping", "Pong", "Ping", "Pong", "Ping", "Pong"},
		},
		{
			name:         "test_all_rate_limited",
			requests:     []string{"Ping", "Pong", "Ping", "Pong", "Ping", "Pong", "Ping", "Pong", "Ping", "Pong"},
			minDuration:  550 * time.Millisecond,
			maxDuration:  750 * time.Millisecond,
			expectedResp: []string{"Pong", "Ping", "Pong", "Ping", "Pong", "Ping", "Pong", "Ping", "Pong", "Ping"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			start := time.Now()
			PingPong(tt.requests)
			elapsed := time.Since(start)

			// Restore stdout
			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Check duration
			if elapsed < tt.minDuration {
				t.Errorf("expected duration at least %v, got %v", tt.minDuration, elapsed)
			}
			if elapsed > tt.maxDuration {
				t.Errorf("expected duration at most %v, got %v", tt.maxDuration, elapsed)
			}

			// Check output mapping
			lines := strings.Split(strings.TrimSpace(output), "\n")
			if len(lines) != len(tt.expectedResp) {
				t.Errorf("expected %d responses, got %d", len(tt.expectedResp), len(lines))
				return
			}

			for i, line := range lines {
				parts := strings.Fields(line)
				if len(parts) < 1 {
					t.Errorf("line %d: invalid format", i+1)
					continue
				}
				response := parts[0]
				if response != tt.expectedResp[i] {
					t.Errorf("line %d: expected %s, got %s", i+1, tt.expectedResp[i], response)
				}
			}

			// Check burst timing (first 4 should have similar timestamps)
			if len(lines) >= 4 {
				timestamps := make([]string, 4)
				for i := 0; i < 4; i++ {
					parts := strings.Fields(lines[i])
					if len(parts) >= 2 {
						timestamps[i] = parts[1]
					}
				}

				// Parse first timestamp
				firstTime := timestamps[0]
				sameCount := 0
				for _, ts := range timestamps {
					if ts == firstTime {
						sameCount++
					}
				}

				// At least 3 out of 4 should be the same (allowing for millisecond boundaries)
				if sameCount < 3 {
					t.Logf("Warning: burst requests may not be processing simultaneously (only %d/4 had same timestamp)", sameCount)
				}
			}
		})
	}
}
