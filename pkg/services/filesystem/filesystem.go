package filesystem

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Hash will take a file and return the hash of the filename.
func Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Extract the file extension from the original file (optional)
	fileExt := filepath.Ext(filePath)
	uniqueFilename := hex.EncodeToString(hash.Sum(nil)) + fileExt

	return uniqueFilename, nil
}

// CheckFile checks if a file exists at the given path and returns the file name or an error.
func CheckFile(filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist: %s", filePath)
		}
		return "", fmt.Errorf("error checking file: %w", err)
	}

	return fileInfo.Name(), nil
}
