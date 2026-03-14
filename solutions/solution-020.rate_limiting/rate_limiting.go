package rate_limiting

import (
	"fmt"
	"time"
)

// TODO: Implement PingPong
// Read README.md for the instructions

func PingPong(requests []string) {

	// 1. Create a buffered channel to act as the rate limiter with burst capacity of 4
	limiter := make(chan time.Time, 4)

	// 2. Pre-fill the limiter with 4 tokens so the first 4 requests can proceed immediately
	for i := 0; i < 4; i++ {
		limiter <- time.Now()
	}

	// 3. Start a goroutine that continuously adds tokens every 100ms
	// This enforces the rate limit of 10 requests per second
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for t := range ticker.C {
			limiter <- t
		}
	}()

	// 4. Create a buffered channel to hold incoming requests
	reqChan := make(chan string, len(requests))

	// 5. Push all requests into the channel
	for _, r := range requests {
		reqChan <- r
	}

	// 6. Close the channel since no more requests will be sent
	close(reqChan)

	// 7. Process requests one by one
	for req := range reqChan {

		// 8. Acquire a token from the limiter before processing
		<-limiter

		// 9. Capture the current timestamp in required format
		timestamp := time.Now().Format("15:04:05.000")

		// 10. Print the opposite response with timestamp
		if req == "Ping" {
			fmt.Printf("Pong %s\n", timestamp)
		} else if req == "Pong" {
			fmt.Printf("Ping %s\n", timestamp)
		}
	}
}
