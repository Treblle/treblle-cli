package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/treblle/treblle-cli/pkg/services/filesystem"
)

// SendToS3 will take a filepath as well as the AWS credentials and upload
// the file to AWS S3 returning the URL of the upload or an error.
func SendToS3(filePath string, accessKeyID string, secretAccessKey string, region string, bucketName string) (string, error) {
	// Generate a unique filename
	uniqueFilename, err := filesystem.Hash(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to generate unique filename: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a session (ensure you have imported the AWS session package)
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		fmt.Printf("Failed to create the AWS Session: %v\n\n", err)
		os.Exit(1)
	}

	// Define the uploader
	uploader := s3manager.NewUploader(sess)

	// Upload the file with the unique filename
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(uniqueFilename),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
