package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

// executeCommand is a helper function to execute a Cobra command
func executeCommand(cmd *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf) // Redirects the command output to buf
	cmd.SetErr(buf) // Redirects the command error output to buf
	cmd.SetArgs(args)

	err := cmd.Execute()
	return buf.String(), err
}

// TestInsightsCmd tests the 'insights' command
func TestInsightsCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		wantOut string
	}{
		{
			name:    "valid args",
			args:    []string{"insights", "dummy/filepath", "--mode", "ide"},
			wantErr: false,
			wantOut: "Processing file: dummy/filepath with mode: ide\n",
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := executeCommand(rootCmd, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("executeCommand() got = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
