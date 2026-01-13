package selecttimeout

import "fmt"

func FunctionOrdered() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)
	c5 := make(chan string)
	// TODO:Implement go-routines
	// Read README.md for the instructions
	go func() {

	}()
	go func() {

	}()
	go func() {

	}()
	go func() {

	}()
	go func() {

	}()
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg3 := <-c3:
			fmt.Println(msg3)
		case msg4 := <-c4:
			fmt.Println(msg4)
		case msg5 := <-c5:
			fmt.Println(msg5)
		}
	}

}
