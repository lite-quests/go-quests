package logging

import (
	"log/slog"
)

// TODO: Implement ProcessSearch
// Read README.md for the instructions

func ProcessSearch(logger *slog.Logger, reqID string, query string, userID int) {
	// 1. Create child logger with request_id attached
	child := logger.With("request_id", reqID)

	// 2. Log processing start
	child.Debug("processing search request")

	// 3. Validate query
	if query == "" {
		child.Warn("empty search query")
		return
	}

	// 4. Validate userID
	if userID < 0 {
		child.Error("invalid user ID")
		return
	}

	// 5. Log valid search
	child.Info(
		"search executed",
		"query", query,
		slog.Group("user", slog.Int("id", userID)),
	)
}
