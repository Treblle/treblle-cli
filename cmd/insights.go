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

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/treblle/treblle-cli/pkg/routes"
	"github.com/treblle/treblle-cli/pkg/types"
	"github.com/treblle/treblle-cli/pkg/views"
	"gopkg.in/yaml.v2"
)

var insightsCmd = &cobra.Command{
	Use:   "insights [file path]",
	Short: "Generate an API Insights report.",
	Long:  "Generate an API Insights report from your OpenAPI specification.",
	Args:  cobra.ExactArgs(1),
	Run:   uploadFile,
}

func init() {
	insightsCmd.Flags().StringP("details", "d", "", "Show the details of a report, options: 'all', 'design', 'performance', 'security'")
	insightsCmd.Flags().BoolP("technology", "t", false, "Show the discovered technology.")
}

// uploadFile is the function that gets called when the Cobra command is executed
func uploadFile(cmd *cobra.Command, args []string) {

	details, _ := cmd.Flags().GetString("details")
	technology, _ := cmd.Flags().GetBool("technology")

	newHeader := pterm.HeaderPrinter{
		Margin:          20,
	}
	newHeader.Println("API Insights")

	filePath := args[0]

	spinnerFile, _ := pterm.DefaultSpinner.Start("Validating OpenAPI Specification.")
	if !checkMime(filePath) {
		spinnerFile.Fail("Cannot use this file, only JSON or YAML is supported")
		os.Exit(1)
	}
	spinnerFile.Success("File validated!")

	spinnerOpen, _ := pterm.DefaultSpinner.Start("Processing OpenAPI Specification.")
	file, err := os.Open(filePath)
	if err != nil {
		spinnerOpen.Fail(fmt.Printf("Failed to open file: %v\n\n", err))
		os.Exit(1)
	}
	defer file.Close()
	content, err := os.ReadFile(filePath)
	if err != nil {
		spinnerOpen.Fail(fmt.Printf("Failed to read file contents: %v\n\n", err))
		os.Exit(1)
	}
	if !validateFile(content, filePath) {
		spinnerOpen.Fail("Could not validate the contents of this file as either YAML or JSON")
		os.Exit(1)
	}
	spinnerOpen.Success("File Processed")

	spinnerRequest, _ := pterm.DefaultSpinner.Start("Sending Request ...")

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		spinnerRequest.Fail(fmt.Printf("Failed to create Form Data: %v\n\n", err))

		os.Exit(1)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		spinnerRequest.Fail(fmt.Printf("Failed to write part from file: %v\n\n", err))
		os.Exit(1)
	}

	writer.Close()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, routes.InsightsUrl(), &requestBody)
	if err != nil {
		spinnerRequest.Fail(fmt.Printf("Failed to create request: %v\n\n", err))
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
		spinnerRequest.Fail(fmt.Printf("Error reading API response: %v\n\n", err))
		os.Exit(1)
	}

	spinnerRequest.Success("Response Received.")

	var apiResponse types.ApiResponse
	err = json.Unmarshal([]byte(body), &apiResponse)
	if err != nil {
		spinnerRequest.Fail(fmt.Printf("Failed to process API Response: %v\n\n", err))
		os.Exit(1)
	}

	views.ShowInsightsDetails(&apiResponse)

	switch details {
	case "performance", "p":
		views.NewInsightsPerformanceView(&apiResponse)
	case "design", "d":
		views.NewInsightsDesignView(&apiResponse)
	case "security", "s":
		views.NewInsightsSecurityView(&apiResponse)
	case "all", "a":
		views.NewInsightsFullView(&apiResponse)
	default:
		views.NewApiInsightsView(&apiResponse)
	}

	if technology {
		views.ShowInsightsTechnologyDiscovery(&apiResponse)
	}
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
