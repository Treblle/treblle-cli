package styles

import "github.com/charmbracelet/lipgloss"

var Text = lipgloss.NewStyle().
	Foreground(lipgloss.CompleteAdaptiveColor{
		Light: lipgloss.CompleteColor{
			TrueColor: "#d7ffae",
			ANSI256:   "193",
			ANSI:      "11",
		},
		Dark: lipgloss.CompleteColor{
			TrueColor: "#d75fee",
			ANSI256:   "163",
			ANSI:      "5",
		},
	})

var Heading = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("6")).
	PaddingBottom(1).
	PaddingTop(1)

var SubHeading = lipgloss.NewStyle().
	Foreground(lipgloss.Color("2")).
	Italic(true)

var Body = lipgloss.NewStyle().
	Foreground(lipgloss.Color("7"))

var Link = lipgloss.NewStyle().
	Foreground(lipgloss.Color("4")).
	Underline(true)

var Key = lipgloss.NewStyle().
	Foreground(lipgloss.Color("1")).
	Bold(true)

var Value = lipgloss.NewStyle().
	Foreground(lipgloss.Color("5"))

var Title = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("3"))

var Divider = lipgloss.NewStyle().
	Foreground(lipgloss.Color("8"))

	// Define a map for grade to color mapping
var ReportsSvgGradeColorMap = map[string]string{
	"A": "#54FF77",
	"B": "#04E8B0",
	"C": "#FFF586",
	"D": "#FFA463",
	"F": "#FF0057",
}

// Define a fallback color constant
const ReportsSvgColorBoundaryFallbackColor = "#6C7285"

var reportsGradeToScorePercentageMap = map[string][2]int{
	"A": {90, 100},
	"B": {80, 89},
	"C": {70, 79},
	"D": {60, 69},
	"F": {0, 59},
}

// Function to map an integer score to a letter grade
func gradeFromScore(score int) string {
	for grade, rangeVal := range reportsGradeToScorePercentageMap {
		if score >= rangeVal[0] && score <= rangeVal[1] {
			return grade
		}
	}
	return "" // Return empty string if no grade matches
}

// Define a map for grade to score percentage range
// Function to get a lipgloss style based on the grade
func StyleForScore(score int) lipgloss.Style {
	color, exists := ReportsSvgGradeColorMap[gradeFromScore(score)]
	if !exists {
		color = ReportsSvgColorBoundaryFallbackColor
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}

// Define a map for grade to score percentage range
// Function to get a lipgloss style based on the grade
func StyleForGrade(grade string) lipgloss.Style {
	color, exists := ReportsSvgGradeColorMap[grade]
	if !exists {
		color = ReportsSvgColorBoundaryFallbackColor
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}
