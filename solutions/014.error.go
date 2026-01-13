package solutions

import (
	"errors"
	"fmt"
	"strings"
)

// Part 1: Define sentinel errors
var (
	// TODO: Define ErrEmptyFilename
	ErrEmptyFilename = errors.New("filename cannot be empty")
	// TODO: Define ErrFileTooLarge
	ErrFileTooLarge = errors.New("file size exceeds limit")
)

// Part 2: Define custom error type
type ValidationError struct {
	Filename string
	Reason   string
}

// TODO: Implement Error() method
func (ve *ValidationError) Error() string {
	return "validation failed for " + ve.Filename + ": " + ve.Reason
}

// Part 3: Validation functions

func ValidateFilename(filename string) error {
	// TODO: Implement
	if filename == "" {
		return ErrEmptyFilename
	}
	return nil
}

func ValidateFileSize(size int64, maxSize int64) error {
	// TODO: Implement
	if size > maxSize {
		return ErrFileTooLarge
	}
	if size < 0 {
		return errors.New("file size cannot be negative")
	}
	return nil
}

func ValidateFileExtension(filename string, allowedExts []string) error {
	// TODO: Implement
	for _, ext := range allowedExts {
		if strings.HasSuffix(filename, ext) {
			return nil
		}
	}
	ve := ValidationError{
		Filename: filename,
		Reason:   "unsupported file extension",
	}
	return &ve
}

func ValidateFile(filename string, size int64, maxSize int64) error {
	// TODO: Implement
	err := ValidateFilename(filename)
	if err != nil {
		return fmt.Errorf("file validation failed: %w", err)
	}
	err = ValidateFileSize(size, maxSize)
	if err != nil {
		return fmt.Errorf("file validation failed: %w", err)
	}
	return nil
}

// Part 4: Error checking

func CanRetry(err error) bool {
	// TODO: Implement
	if err == nil {
		return false
	}
	var ve *ValidationError
	if errors.As(err, &ve) {
		return true
	}
	return false
}
