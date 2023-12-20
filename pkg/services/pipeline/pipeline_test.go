package pipeline

import (
	"fmt"
	"testing"
)

// mockHandler doubles the input if it's an integer
type mockHandler struct{}

func (h *mockHandler) Process(input interface{}) (interface{}, error) {
	num, ok := input.(int)
	if !ok {
		return nil, fmt.Errorf("not an integer")
	}
	return num * 2, nil
}

// mockErrorHandler always returns an error
type mockErrorHandler struct{}

func (h *mockErrorHandler) Process(input interface{}) (interface{}, error) {
	return nil, fmt.Errorf("error occurred")
}

// TestPipelineSuccess tests the pipeline with successful processing
func TestPipelineSuccess(t *testing.T) {
	p := NewPipeline(&mockHandler{})
	p.Start()

	go func() {
		p.Input() <- 10
		close(p.Input())
	}()

	for result := range p.Output() {
		if result != 20 {
			t.Errorf("Expected 20, got %v", result)
		}
	}

	if err := <-p.Errors(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

// TestPipelineError tests the pipeline with an error
func TestPipelineError(t *testing.T) {
	p := NewPipeline(&mockErrorHandler{})
	p.Start()

	go func() {
		p.Input() <- 10
		close(p.Input())
	}()

	select {
	case <-p.Output():
		t.Error("Expected no output, got some")
	case err := <-p.Errors():
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	}
}
