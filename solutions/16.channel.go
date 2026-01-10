package solutions

// TODO:Implement ReadFromBoth function
// Read README.md for the instructions
func ReadFromBoth(ch1 chan string, ch2 chan string) string {
	val1 := <-ch1
	val2 := <-ch2
	return "read: " + val1 + " & " + val2
}

// TODO:Implement ReadFromBoth function
// Read README.md for the instructions
func WriteToBoth(ch1 chan string, ch2 chan string, msg string) {
	go func() {
		formatted := "write: " + msg
		ch1 <- formatted
		ch2 <- formatted
	}()
}

// TODO:Implement ReadFromBoth function
// Read README.md for the instructions
func ReadThenWrite(input chan string, output chan string) {
	val := <-input
	output <- "transform: " + val
}
