package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Send(url string, filePath string) ([]byte, error) {
	// Create a JSON payload
	payload := map[string]string{"file_name": filePath}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	token := os.Getenv("API_INSIGHTS_TOKEN")

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+token)

	ext := filepath.Ext(filePath)
	idempotencyKey := filePath[:len(filePath)-len(ext)]

	req.Header.Set("Idempotency-Key", idempotencyKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
