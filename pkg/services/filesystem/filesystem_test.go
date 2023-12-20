package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGetUniqueFilename tests the getUniqueFilename function
func TestHash(t *testing.T) {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write some content to the file
	content := []byte("Hello, world!")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %s", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %s", err)
	}

	// Test getUniqueFilename
	uniqueFilename, err := Hash(tmpfile.Name())
	if err != nil {
		t.Errorf("getUniqueFilename returned an error: %s", err)
	}
	if uniqueFilename == "" {
		t.Error("Expected a unique filename, got an empty string")
	}

	// Optionally, you can add more assertions here if you want
	// to test specific aspects of the unique filename.
}

// TestCheckFile tests the CheckFile function.
func TestCheckFile(t *testing.T) {
	// Create a temporary file to test
	tempFile, err := os.CreateTemp("", "checkfiletest.*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Get the base name of the file
	baseName := filepath.Base(tempFile.Name())

	// Test with existing file
	filename, err := CheckFile(tempFile.Name())
	if err != nil {
		t.Errorf("CheckFile returned an error for an existing file: %s", err)
	}
	if filename != baseName {
		t.Errorf("Expected filename %s, got %s", baseName, filename)
	}

	// Test with non-existing file
	_, err = CheckFile("/path/to/non/existing/file.txt")
	if err == nil {
		t.Error("Expected an error for a non-existing file, got nil")
	}
}
