package files

import (
	"errors"
	"os"
	"strings"
)

// TODO: Implement ProcessFile
// Read README.md for the instructions

func ProcessFile() error {
	// 1. Read the file contents
	data, err := os.ReadFile("task.txt")
	if err != nil {
		return err
	}

	// 2. Split into words
	words := strings.Fields(string(data))

	// 3. Ensure there are at least 10 words
	if len(words) < 10 {
		return errors.New("task.txt must contain at least 10 words")
	}

	// 4. Capitalize first 5 words
	for i := 0; i < 5; i++ {
		words[i] = strings.ToUpper(words[i])
	}

	// 5. Lowercase last 5 words
	for i := len(words) - 5; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	// 6. Join words with a single space and add newline
	result := strings.Join(words, " ") + "\n"

	// 7. Write back to the file (overwrite existing content)
	err = os.WriteFile("task.txt", []byte(result), 0644)
	if err != nil {
		return err
	}

	// 8. Return nil on success
	return nil
}
