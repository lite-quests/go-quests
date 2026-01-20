package error

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the error Quest 🎉", colorReset)
	}
	os.Exit(code)
}

func TestValidateFilename(t *testing.T) {
	expectedErr := errors.New("filename cannot be empty")

	err := ValidateFilename("")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected %q, got %q", expectedErr.Error(), err.Error())
	}

	err = ValidateFilename("file.txt")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestValidateFileSize(t *testing.T) {
	tests := []struct {
		name    string
		size    int64
		maxSize int64
		wantErr string
	}{
		{
			name:    "negative size",
			size:    -1,
			maxSize: 100,
			wantErr: "file size cannot be negative",
		},
		{
			name:    "too large",
			size:    200,
			maxSize: 100,
			wantErr: "file size exceeds limit",
		},
		{
			name:    "valid size",
			size:    50,
			maxSize: 100,
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFileSize(tt.size, tt.maxSize)

			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("expected nil, got %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if err.Error() != tt.wantErr {
				t.Fatalf("expected %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}

func TestValidateFileExtension(t *testing.T) {
	err := ValidateFileExtension("doc.pdf", []string{".txt", ".md"})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected ValidationError, got %T", err)
	}

	if ve.Filename != "doc.pdf" {
		t.Fatalf("expected filename doc.pdf, got %s", ve.Filename)
	}

	if ve.Reason != "unsupported file extension" {
		t.Fatalf("unexpected reason: %s", ve.Reason)
	}
}

func TestValidateFileExtensionAllowed(t *testing.T) {
	err := ValidateFileExtension("readme.md", []string{".txt", ".md"})
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestValidateFileWrapping(t *testing.T) {
	err := ValidateFile("", 10, 100)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, ErrEmptyFilename) {
		t.Fatalf("expected wrapped ErrEmptyFilename, got %v", err)
	}
}

func TestCanRetry(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "empty filename",
			err:  fmt.Errorf("wrapped: %w", errors.New("filename cannot be empty")),
			want: false,
		},
		{
			name: "file too large",
			err:  fmt.Errorf("wrapped: %w", errors.New("file size exceeds limit")),
			want: false,
		},
		{
			name: "validation error",
			err: &ValidationError{
				Filename: "doc.pdf",
				Reason:   "unsupported file extension",
			},
			want: true,
		},
		{
			name: "unknown error",
			err:  errors.New("random error"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanRetry(tt.err)
			if got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
