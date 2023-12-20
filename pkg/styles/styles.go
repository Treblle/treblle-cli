package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var Text = lipgloss.NewStyle().
	PaddingLeft(4).
	PaddingRight(4).
	PaddingTop(4).
	PaddingBottom(4)

var Info = Text.
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4"))

var Success = Text.
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4"))
