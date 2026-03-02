package context_quest

import (
	"context"
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
}
