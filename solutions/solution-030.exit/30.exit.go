package solutions

import (
	"fmt"
	"os"
)

func ExitWithStatus() {
	defer fmt.Print("should not run")
	os.Exit(3)
}
