package error

// Part 1: Define sentinel errors
var (
// TODO: Define ErrEmptyFilename
// TODO: Define ErrFileTooLarge
)

// Part 2: Define custom error type
type ValidationError struct {
	Filename string
	Reason   string
}

// TODO: Implement Error() method
func (ve *ValidationError) Error() string {
	return ""
}

// Part 3: Validation functions

func ValidateFilename(filename string) error {
	// TODO: Implement
	return nil
}

func ValidateFileSize(size int64, maxSize int64) error {
	// TODO: Implement
	return nil
}

func ValidateFileExtension(filename string, allowedExts []string) error {
	// TODO: Implement
	return nil
}

func ValidateFile(filename string, size int64, maxSize int64) error {
	// TODO: Implement
	return nil
}

// Part 4: Error checking

func CanRetry(err error) bool {
	// TODO: Implement
	return false
}
