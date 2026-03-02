package context_quest

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRunSmartDownloader(t *testing.T) {
	api := &CDNAPI{}

	RunSmartDownloader(api)

	if api.LogCtx == nil {
		t.Fatal("LogActivity was not called or passed a nil context")
	}
	val := api.LogCtx.Value(TraceKey("trace_id"))
	if val == nil {
		t.Errorf("expected LogActivity context to have a value for TraceKey(\"trace_id\")")
	} else if valStr, ok := val.(string); !ok || valStr != "smart-dl-123" {
		t.Errorf("expected LogActivity context to have value 'smart-dl-123', got %v", val)
	}

	if api.VerifyCtx == nil {
		t.Fatal("VerifyChecksum was not called or passed a nil context")
	}
	if !strings.Contains(strings.ToLower(fmt.Sprintf("%v", api.VerifyCtx)), "todo") {
		t.Errorf("expected VerifyChecksum to receive context.TODO(), got %T (%v)", api.VerifyCtx, api.VerifyCtx)
	}

	if api.DNSCtx == nil {
		t.Fatal("DNSLookup was not called or passed a nil context")
	}
	deadline, ok := api.DNSCtx.Deadline()
	if !ok {
		t.Errorf("expected DNSLookup context to have a deadline (WithTimeout)")
	} else {

		dur := time.Until(deadline)
		if dur > 100*time.Millisecond || dur < 0 {
			t.Errorf("expected timeout of ~100ms, but deadline is in %v", dur)
		}
	}

	if api.DownloadCtx == nil {
		t.Fatal("DownloadChunks was not called or passed a nil context")
	}
	select {
	case <-api.DownloadCtx.Done():

		err := api.DownloadCtx.Err()
		if err != context.Canceled {
			t.Errorf("expected DownloadCtx to fail with context.Canceled, got %v", err)
		}
	default:
		t.Errorf("expected DownloadCtx to be already cancelled, but Done() channel is open. Did you call cancel() BEFORE passing it?")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the context Quest 🎉", colorReset)
	}
	os.Exit(code)
}
