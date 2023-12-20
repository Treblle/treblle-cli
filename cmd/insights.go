package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/treblle/treblle-cli/pkg/http"
	"github.com/treblle/treblle-cli/pkg/styles"
)

// insightsCmd is a command to send your OpenAPI Specification to API Insights.
var insightsCmd = &cobra.Command{
	Use:   "insights [file]",
	Short: "Run API Insights on your OpenAPI Specification.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]

		// Check if file exists
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			fmt.Fprintf(cmd.OutOrStderr(), "File does not exist: %s\n", filepath)
			os.Exit(1)
		}

		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error reading mode flag: %s\n", err)
			os.Exit(1)
		}

		fmt.Fprintln(cmd.OutOrStdout(), styles.Info.Render("Processing file: "+filepath+" using mode: "+mode))

		_, err = http.Send("https://httpdump.app/dumps/35be3fed-6393-4a21-89dc-0c4a5ebb59e0", filepath)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Request Failed: %s\n\n", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(insightsCmd)

	insightsCmd.Flags().String("mode", "CLI", "Mode of operation (CLI, IDE, or CI)")
}
