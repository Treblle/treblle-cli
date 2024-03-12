package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "treblle-cli",
	Version: "1.0.0",
	Short:   "The Treblle CLI tool.",
}

// init will initialize the CLI, adding all sub commands.
func init() {
	rootCmd.AddCommand(insightsCmd)
}

// Execute the CLI tool.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Something went wrong: %v\n\n", err)
		os.Exit(1)
	}
}
