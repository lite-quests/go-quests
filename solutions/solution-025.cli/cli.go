package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Read README.md for the instructions

func RunCLI(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("No subcommand provided")
	}

	switch args[1] {
	case "count":
		countCmd := flag.NewFlagSet("count", flag.ContinueOnError)
		if err := countCmd.Parse(args[2:]); err != nil {
			return err
		}
		if countCmd.NArg() < 1 {
			return fmt.Errorf("The file to count is not provided")
		}
		filename := countCmd.Arg(0)
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("The file to count doesn't exist")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lcnt := 0
		for scanner.Scan() {
			lcnt++
		}
		fmt.Printf("lines: %d\n", lcnt)

	case "search":
		searchCmd := flag.NewFlagSet("search", flag.ContinueOnError)
		caseInsensitive := searchCmd.Bool("case-insensitive", false, "Ignore case")
		if err := searchCmd.Parse(args[2:]); err != nil {
			return err
		}
		if searchCmd.NArg() < 2 {
			return fmt.Errorf("The file to search is not provided")
		}
		pattern := searchCmd.Arg(0)
		filename := searchCmd.Arg(1)

		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("The file to search doesn't exist")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		if *caseInsensitive {
			pattern = strings.ToLower(pattern)
		}
		for scanner.Scan() {
			line := scanner.Text()
			checkLine := line
			if *caseInsensitive {
				checkLine = strings.ToLower(checkLine)
			}
			if strings.Contains(checkLine, pattern) {
				fmt.Println(line)
			}
		}

	default:
		return fmt.Errorf("Unknown subcommand: %s", args[1])
	}
	return nil
}
