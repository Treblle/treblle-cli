package apiinsights

import "github.com/aws/aws-sdk-go/aws/credentials"

type PipelineInput struct {
	Filepath       string
	AWSCredentials *credentials.Credentials
}
