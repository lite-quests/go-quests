package solutions

import (
	"context"
	"time"
)

// Read README.md for the instructions

// TraceKey is a custom type to prevent context key collisions
type TraceKey string

// CDNAPI is just a stub for testing.
type CDNAPI struct {
	LogCtx      context.Context
	DNSCtx      context.Context
	DownloadCtx context.Context
	VerifyCtx   context.Context
}

// These are already implemented for you by the test suite.
func (api *CDNAPI) LogActivity(ctx context.Context)    { api.LogCtx = ctx }
func (api *CDNAPI) DNSLookup(ctx context.Context)      { api.DNSCtx = ctx }
func (api *CDNAPI) DownloadChunks(ctx context.Context) { api.DownloadCtx = ctx }
func (api *CDNAPI) VerifyChecksum(ctx context.Context) { api.VerifyCtx = ctx }

// TODO: Implement RunSmartDownloader
func RunSmartDownloader(api *CDNAPI) {
	// 1. Root context
	rootCtx := context.Background()

	// 2. Tracing (LogActivity)
	traceCtx := context.WithValue(rootCtx, TraceKey("trace_id"), "smart-dl-123")
	api.LogActivity(traceCtx)

	// 3. DNS Lookup with timeout (100ms)
	dnsCtx, dnsCancel := context.WithTimeout(rootCtx, 100*time.Millisecond)
	defer dnsCancel()
	api.DNSLookup(dnsCtx)

	// 4. Concurrent Download with cancel
	downloadCtx, cancel := context.WithCancel(rootCtx)
	cancel() // simulate early completion
	api.DownloadChunks(downloadCtx)

	// 5. Verify Checksum with TODO context
	api.VerifyChecksum(context.TODO())

}
