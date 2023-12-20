package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	apiinsights "github.com/treblle/treblle-cli/pkg/services/api-insights"
	"github.com/treblle/treblle-cli/pkg/services/pipeline"
)

// insightsCmd is a command to send your OpenAPI Specification to API Insights.
var insightsCmd = &cobra.Command{
	Use:   "insights [file]",
	Short: "Run API Insights on your OpenAPI Specification.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		filepath := args[0]

		// Initialize success and error channels
		successChan := make(chan string)
		errorChan := make(chan error)

		// Start running the pipeline
		go runInsightsCmd(filepath, successChan, errorChan)

		// Initialize the Bubble Tea program
		p := tea.NewProgram(initialModel())

		// Run the Bubble Tea program in a separate goroutine
		go func() {
			if _, err := p.Run(); err != nil {
				fmt.Println("Error running Bubble Tea program:", err)
			}
		}()

		// Handle pipeline responses
		select {
		case errMsg := <-errorChan:
			fmt.Printf("Error: %v\n\n", errMsg)
			p.Send(errMsg) // Send error message to Bubble Tea program
		case successMsg := <-successChan:
			p.Send(successMsg) // Send success message to Bubble Tea program
		}
	},
}

// init will initialize the Insights Command.
func init() {
	insightsCmd.Flags().String("mode", "CLI", "Mode of operation (CLI, IDE, or CI)")
}

type errMsg error

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Analysing Your OpenAPI Specification.\n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}

	return str
}

func runInsightsCmd(path string, successChan chan<- string, errorChan chan<- error) {
	// Artificial wait
	time.Sleep(1 * time.Second) // Wait for 5 seconds

	// Resolve the absolute file path
	absPath, err := filepath.Abs(path)
	if err != nil {
		errorChan <- fmt.Errorf("failed to get absolute file path: %w", err)
		return
	}

	pipeline := pipeline.NewPipeline(
		apiinsights.CheckFileHandler{},
		apiinsights.UploadToS3Handler{},
		apiinsights.SendAPIRequestHandler{},
	)

	pipeline.Start()

	// get absolute file path?

	go func() {
		defer close(pipeline.Input())
		pipeline.Input() <- absPath
	}()

	go func() {
		for result := range pipeline.Output() {
			successChan <- result.(string)
		}
	}()

	go func() {
		for err := range pipeline.Errors() {
			errorChan <- err
		}
	}()
}
