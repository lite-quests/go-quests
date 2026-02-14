package ratelimiting

// PingPong processes requests with rate limiting.
//
// Rate limiting rules:
// - First 4 requests: processed immediately (burst)
// - Subsequent requests: max 10 per second (one every 100ms)
//
// Response mapping:
// - "Ping" → "Pong <timestamp>"
// - "Pong" → "Ping <timestamp>"
//
// Implementation hints:
// - Create a buffered channel of time.Time (size 4) for the rate limiter
// - Pre-fill it with 4 tokens using a loop
// - Launch a goroutine that sends tokens via time.Tick every 100ms
// - Create another buffered channel to hold all requests
// - Range over the requests channel, acquiring a token before each print
func PingPong(requests []string) {
	// TODO: Implement rate limiting with burst capacity
	// NOTE: Before Implmenting consider going through the following resources:
	// - https://gobyexample.com/rate-limiting
}
