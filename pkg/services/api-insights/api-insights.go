package apiinsights

import (
	"fmt"
	"log"
	"os"

	"github.com/treblle/treblle-cli/pkg/services/aws"
	"github.com/treblle/treblle-cli/pkg/services/filesystem"
	"github.com/treblle/treblle-cli/pkg/services/http"
)

// These are dummy implementations. Replace with actual logic.
type CheckFileHandler struct{}

func (h CheckFileHandler) Process(input interface{}) (interface{}, error) {
	filePath, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("input is not a valid string")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	log.Printf("File exists: %s", filePath)
	return filePath, nil
}

type UploadToS3Handler struct{}

func (h UploadToS3Handler) Process(input interface{}) (interface{}, error) {
	filePath, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("input is not a valid string")
	}

	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_BUCKET_NAME")

	// Implement the S3 upload logic. Assume it returns the URL of the uploaded file
	uploadedURL, err := aws.SendToS3(filePath, accessKeyID, secretAccessKey, region, bucketName)
	if err != nil {
		return nil, err
	}

	log.Printf("Uploaded to S3: %s", uploadedURL)
	return filePath, nil
}

type SendAPIRequestHandler struct{}

func (h SendAPIRequestHandler) Process(input interface{}) (interface{}, error) {
	uniqueFilename, err := filesystem.Hash(input.(string))
	if err != nil {
		return "", fmt.Errorf("failed to generate unique filename: %w", err)
	}

	// Send the API request. Implement this according to your API logic.
	response, err := http.Send("https://staging-api.apiinsights.io/v1/reports", uniqueFilename)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}

	return string(response), nil
}
