package main

import (
	"errors"
	"io/fs"
	"os"
	"strings"
	"testing"
)

func TestReadFile_Success(t *testing.T) {
	t.Run("Reads valid file content accurately", func(t *testing.T) {
		// Create a temporary file for the test
		content := "Hello, Go testing world!"
		tmpFile, err := os.CreateTemp("", "test_read_file_*.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		// Clean up the temp file after the test ends
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(content); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		tmpFile.Close()

		// Execute the target function
		got, err := ReadFile(tmpFile.Name())
		if err != nil {
			t.Fatalf("ReadFile returned an unexpected error: %v", err)
		}

		if got != content {
			t.Errorf("ReadFile() = %q; want %q", got, content)
		}
	})
}

func TestReadFile_FailureAndErrorWrapping(t *testing.T) {
	t.Run("Wraps OS error and supports errors.Is checks", func(t *testing.T) {
		invalidPath := "this_file_does_not_exist_12345.txt"

		// Execute target function on non-existent path
		_, err := ReadFile(invalidPath)
		if err == nil {
			t.Fatal("Expected an error for an invalid path, but got nil")
		}

		// 1. Verify specific text wrapping format requirement
		expectedPrefix := "ReadFile " + invalidPath + ":"
		if !strings.HasPrefix(err.Error(), expectedPrefix) {
			t.Errorf("Error string %q does not match required wrapping format prefix %q", err.Error(), expectedPrefix)
		}

		// 2. Verify %w wrapper works perfectly by testing unwrapping with errors.Is
		// A non-existent file triggers an underlying fs.ErrNotExist error
		if !errors.Is(err, fs.ErrNotExist) {
			t.Errorf("Error wrapping failed: errors.Is could not find underlying fs.ErrNotExist in: %v", err)
		}
	})
}

func TestReadFile_MainWorkflowLogic(t *testing.T) {
	t.Run("Validation of errors.As usability on path errors", func(t *testing.T) {
		invalidPath := "missing_file_for_as_test.txt"
		_, err := ReadFile(invalidPath)

		// Verify the underlying OS error can still be extracted using errors.As
		var pathErr *fs.PathError
		if !errors.As(err, &pathErr) {
			t.Error("Expected error to wrap an underlying *fs.PathError, but errors.As failed")
		} else if pathErr.Path != invalidPath {
			t.Errorf("Extracted PathError path = %q; want %q", pathErr.Path, invalidPath)
		}
	})
}
