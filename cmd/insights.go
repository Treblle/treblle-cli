package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// insightsCmd is a command to send your OpenAPI Specification to API Insights.
var insightsCmd = &cobra.Command{
	Use:   "insights",
	Short: "Run API Insights on your OpenAPI Specification.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]
		mode, _ := cmd.Flags().GetString("mode")

		fmt.Fprintf(cmd.OutOrStdout(), "Processing file: %s with mode: %s\n", filepath, mode)
	},
}

func init() {
	rootCmd.AddCommand(insightsCmd)

	insightsCmd.PersistentFlags().String("mode", "CLI", "Mode of operation: CLI (default), IDE, or CI")
}
