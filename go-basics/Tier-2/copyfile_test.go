package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile_Success(t *testing.T) {
	t.Run("Successfully copies file contents and matches file sizes", func(t *testing.T) {
		// Set up a temporary directory for clean sandbox execution
		tmpDir, err := os.MkdirTemp("", "copyfile_test_dir_*")
		if err != nil {
			t.Fatalf("Failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		srcPath := filepath.Join(tmpDir, "source.txt")
		dstPath := filepath.Join(tmpDir, "destination.txt")

		// Create source file with explicit payload
		dummyContent := []byte("Validating exact payload content and matching file sizes.")
		err = os.WriteFile(srcPath, dummyContent, 0644)
		if err != nil {
			t.Fatalf("Failed to write initial source file: %v", err)
		}

		// Execute your copy function
		err = CopyFile(srcPath, dstPath)
		if err != nil {
			t.Fatalf("CopyFile returned an unexpected error: %v", err)
		}

		// 1. Confirm destination file exists
		dstInfo, err := os.Stat(dstPath)
		if err != nil {
			t.Fatalf("Failed to look up destination file status: %v", err)
		}

		// 2. Verify source and destination file sizes match exactly (Task Requirement 5)
		srcInfo, err := os.Stat(srcPath)
		if err != nil {
			t.Fatalf("Failed to look up source file status: %v", err)
		}

		if srcInfo.Size() != dstInfo.Size() {
			t.Errorf("File size mismatch! Source is %d bytes, Destination is %d bytes", srcInfo.Size(), dstInfo.Size())
		}

		// 3. Deeper validation to ensure contents are identical
		copiedContent, err := os.ReadFile(dstPath)
		if err != nil {
			t.Fatalf("Failed to read back copied file content: %v", err)
		}

		if !bytes.Equal(dummyContent, copiedContent) {
			t.Error("Data corruption detected: destination file contents do not match source")
		}
	})
}

func TestCopyFile_SourceMissing(t *testing.T) {
	t.Run("Propagates errors immediately if the source file is missing", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copyfile_err_dir_*")
		if err != nil {
			t.Fatalf("Failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		nonExistentSrc := filepath.Join(tmpDir, "does_not_exist.txt")
		dstPath := filepath.Join(tmpDir, "output.txt")

		// Execute with bad source (Step 4 Requirement: return any error)
		err = CopyFile(nonExistentSrc, dstPath)
		if err == nil {
			t.Fatal("Expected an error when attempting to copy a missing source file, but got nil")
		}
	})
}

func TestCopyFile_DestinationInvalid(t *testing.T) {
	t.Run("Propagates errors if destination creation fails", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copyfile_dst_err_dir_*")
		if err != nil {
			t.Fatalf("Failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tmpDir)

		srcPath := filepath.Join(tmpDir, "source.txt")
		err = os.WriteFile(srcPath, []byte("test data"), 0644)
		if err != nil {
			t.Fatalf("Failed to create source file: %v", err)
		}

		// Pointing destination path directly to an impossible directory path
		invalidDst := filepath.Join(tmpDir, "non_existent_folder_abc", "target.txt")

		// Execute (Step 4 Requirement: return any error)
		err = CopyFile(srcPath, invalidDst)
		if err == nil {
			t.Fatal("Expected a file path creation error for missing destination directories, but got nil")
		}
	})
}
