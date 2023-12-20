package cmd

import (
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		wantOut string
	}{
		{
			name:    "help flag",
			args:    []string{"--help"},
			wantErr: false,
			wantOut: "The Treblle CLI tool.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := executeCommand(rootCmd, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(gotOut, tt.wantOut) {
				t.Errorf("RootCommand() output does not contain %q, got: %s", tt.wantOut, gotOut)
			}
		})
	}
}
