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
var ReportsSvgGradeColorMap = map[string]lipgloss.CompleteAdaptiveColor{
	"A": {
		Light: lipgloss.CompleteColor{
			TrueColor: "#2ABC48",
			ANSI256:   "6",
			ANSI:      "6",
		},
		Dark: lipgloss.CompleteColor{
			TrueColor: "#54FF77",
			ANSI256:   "8",
			ANSI:      "8",
		},
	},
	"B": {
		Light: lipgloss.CompleteColor{
			TrueColor: "#04E8B0",
			ANSI256:   "14",
			ANSI:      "14",
		},
		Dark: lipgloss.CompleteColor{
			TrueColor: "#04E8B0",
			ANSI256:   "14",
			ANSI:      "14",
		},
	},
	"C": {
		Light: lipgloss.CompleteColor{
			TrueColor: "#B59D03",
			ANSI256:   "3",
			ANSI:      "3",
		},
		Dark: lipgloss.CompleteColor{
			TrueColor: "#FFF36B",
			ANSI256:   "11",
			ANSI:      "1",
		},
	},
	"D": {
		Light: lipgloss.CompleteColor{
			TrueColor: "#FF862F",
			ANSI256:   "11",
			ANSI:      "1",
		},
		Dark: lipgloss.CompleteColor{
			TrueColor: "#FF862F",
			ANSI256:   "11",
			ANSI:      "11",
		},
	},
	"F": {
		Light: lipgloss.CompleteColor{
			TrueColor: "#FF0057",
			ANSI256:   "9",
			ANSI:      "9",
		},
		Dark: lipgloss.CompleteColor{
			TrueColor: "#FF0057",
			ANSI256:   "9",
			ANSI:      "9",
		},
	},
}

// Define a fallback color constant
var ReportsSvgColorBoundaryFallbackColor = lipgloss.CompleteAdaptiveColor{
	Light: lipgloss.CompleteColor{
		TrueColor: "#3C3C56",
		ANSI256:   "4",
		ANSI:      "4",
	},
	Dark: lipgloss.CompleteColor{
		TrueColor: "#E6EAF6",
		ANSI256:   "15",
		ANSI:      "15",
	},
}

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

	return lipgloss.NewStyle().Foreground(color)
}

// Define a map for grade to score percentage range
// Function to get a lipgloss style based on the grade
func StyleForGrade(grade string) lipgloss.Style {
	color, exists := ReportsSvgGradeColorMap[grade]
	if !exists {
		color = ReportsSvgColorBoundaryFallbackColor
	}

	return lipgloss.NewStyle().Foreground(color)
}
