package solutions

import "os"

func ConnectDB(host string, port int, credentialsValid bool) {

	// Case 1: Host unreachable
	if host == "" {
		os.Exit(1)
	}

	// Case 2: Invalid port
	if port <= 0 || port > 65535 {
		os.Exit(2)
	}

	// Case 3: Authentication failed
	if !credentialsValid {
		os.Exit(3)
	}

	// Case 4: Success
	os.Exit(0)
}
