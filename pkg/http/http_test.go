package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestSend(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// You can add logic here to return different responses based on the request
		if r.Method != http.MethodPost {
			t.Errorf("Expected 'POST' request, got '%s'", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mock response"))
	}))
	defer server.Close()

	// Test cases
	tests := []struct {
		name      string
		filePath  string
		want      string
		wantError bool
	}{
		{
			name:      "Valid JSON",
			filePath:  "testdata/valid.json",
			want:      "mock response",
			wantError: false,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with test data
			file, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(file.Name())

			// Write test data to the file
			if _, err := file.WriteString(`{"test": "data"}`); err != nil {
				t.Fatal(err)
			}
			if err := file.Close(); err != nil {
				t.Fatal(err)
			}

			// Call the function under test
			got, err := Send(server.URL, file.Name())
			if (err != nil) != tt.wantError {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantError)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("Send() = %v, want %v", got, tt.want)
			}
		})
	}
}
