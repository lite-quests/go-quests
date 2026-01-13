package solutions

import (
	"fmt"
	"time"
)

func FunctionOrdered() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)
	c5 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		c1 <- "from c1"
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		c4 <- "from c4"
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		c2 <- "from c2"
	}()

	go func() {
		time.Sleep(40 * time.Millisecond)
		c5 <- "from c5"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		c3 <- "from c3"
	}()
	for i := 0; i < 5; i++ {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		case msg := <-c3:
			fmt.Println(msg)
		case msg := <-c4:
			fmt.Println(msg)
		case msg := <-c5:
			fmt.Println(msg)
		}
	}
}
