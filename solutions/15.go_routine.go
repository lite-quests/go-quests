package solutions

func SendRequest(input string) string {
	// TODO: implement
	// Read README.md for the instructions
	ch := make(chan string)
	go Server(input, ch)
	response := <-ch
	return response
}

func Server(request string, response chan string) {
	// TODO: implement
	// Read README.md for the instructions
	response <- "processed: " + request
}
