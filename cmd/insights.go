package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/treblle/treblle-cli/pkg/routes"
	"gopkg.in/yaml.v2"
)

var insightsCmd = &cobra.Command{
	Use:   "insights [file path]",
	Short: "Upload a file to the server",
	Args:  cobra.ExactArgs(1),
	Run:   uploadFile,
}

// uploadFile is the function that gets called when the Cobra command is executed
func uploadFile(cmd *cobra.Command, args []string) {
	filePath := args[0]

	if !checkMime(filePath) {
		fmt.Println("Cannot use this file, only JSON or YAML is supported")
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file contents: %v\n\n", err)
		os.Exit(1)
	}

	if !validateFile(content, filePath) {
		fmt.Println("Could not validate the contents of this file as either YAML or JSON")
		os.Exit(1)
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		fmt.Printf("Failed to create Form Data: %v\n\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Printf("Failed to write part from file: %v\n\n", err)
		os.Exit(1)
	}

	writer.Close()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, routes.InsightsUrl(), &requestBody)
	if err != nil {
		fmt.Printf("Failed to create request: %v\n\n", err)
		os.Exit(1)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Request Failed: %v\n\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading API response: %v\n\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response Body: %s\n\n", string(body))
	fmt.Printf("Status: %v\n\n", response.Status)

}

func checkMime(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	return ext == ".json" || ext == ".yaml" || ext == ".yml"
}

func validateFile(content []byte, filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".json":
		return json.Unmarshal(content, &json.RawMessage{}) == nil
	case ".yaml", ".yml":
		return yaml.Unmarshal(content, new(interface{})) == nil
	default:
		return false
	}
}
