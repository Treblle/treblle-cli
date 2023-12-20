package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func Send(url string, filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var jsonData []byte

	if json.Valid(fileData) {
		jsonData = fileData
	} else {
		var data interface{}
		if err := yaml.Unmarshal(fileData, &data); err != nil {
			return "", err
		}
		jsonData, err = json.Marshal(data)
		if err != nil {
			return "", err
		}
	}

	// Create the request
	// Create and send the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}
