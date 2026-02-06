package exit_quest

import (
	"fmt"
)

func ExitWithStatus() {
	defer fmt.Print("should not run")
	// above statement must not print
	// write your solution here
}
