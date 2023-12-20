package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "treblle",
	Short: "The Treblle CLI tool.",

	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Something went wrong: %v\n\n", err)
		os.Exit(1)
	}
}
