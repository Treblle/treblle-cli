package views

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/treblle/treblle-cli/pkg/styles"
	"github.com/treblle/treblle-cli/pkg/types"
)

var sectionStyle = pterm.NewStyle(pterm.FgLightWhite, pterm.Bold)

func ShowInsightsDetails(apiResponse *types.ApiResponse) {
	prefixPrinter := pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.InfoMessageStyle,
		Prefix: pterm.Prefix{
			Style: pterm.NewStyle(pterm.FgLightWhite, pterm.BgGray),
			Text:  "INFO",
		},
	}
	pterm.DefaultSection.WithStyle(sectionStyle).Println("Details")

	prefixPrinter.Prefix.Text = "Title"
	prefixPrinter.Println(
		pterm.NewStyle(pterm.FgCyan).Sprint(apiResponse.Report.Title),
	)

	prefixPrinter.Prefix.Text = "Description"
	prefixPrinter.Println(
		pterm.NewStyle(pterm.FgCyan).Sprint(apiResponse.Report.Description),
	)

	prefixPrinter.Prefix.Text = "View in Browser"
	linkStyle := pterm.NewStyle(pterm.FgLightCyan, pterm.Underscore)
	prefixPrinter.Println(
		linkStyle.Sprint(apiResponse.Report.ShareURL),
	)

	prefixPrinter.Prefix.Text = "Base URL"
	prefixPrinter.Println(
		linkStyle.Sprint(apiResponse.Report.BaseURL),
	)

	prefixPrinter.Prefix.Text = "Endpoints"
	prefixPrinter.Println(
		pterm.NewStyle(pterm.FgLightWhite).Sprint(apiResponse.Report.TotalEndpoints),
	)

	prefixPrinter.Prefix.Text = "Overall"
	prefixPrinter.Println(
		pterm.NewStyle(pterm.FgLightWhite).Sprint(apiResponse.Report.ScorePercentage),
	)
}

func ShowInsightsTechnologyDiscovery(apiResponse *types.ApiResponse) {
	testData := pterm.TableData{{"Technologies Discovered"}}
	for _, technology := range apiResponse.Report.Technologies {
		row := []string{
			pterm.NewStyle(pterm.FgLightCyan).Sprint(technology.Name),
		}
		testData = append(testData, row)
	}

	pterm.DefaultTable.WithData(testData).Render()

	fmt.Println(styles.Divider.Render("-----------------------------------"))
}

func NewApiInsightsView(apiResponse *types.ApiResponse) {
	prefixPrinter := pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.InfoMessageStyle,
		Prefix: pterm.Prefix{
			Style: pterm.NewStyle(pterm.FgWhite),
			Text:  "INFO",
		},
	}
	pterm.DefaultSection.WithStyle(sectionStyle).Println("API Insights Report")

	for _, category := range apiResponse.Report.Categories {
		prefixPrinter.Prefix.Text = "Category:"
		prefixPrinter.Println(
			pterm.NewStyle(pterm.FgCyan, pterm.Bold).Sprint(category.Title),
		)

		prefixPrinter.Prefix.Text = "Score:"
		prefixPrinter.Println(
			styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
		)

		prefixPrinter.Prefix.Text = "Grade:"
		prefixPrinter.Println(
			styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
		)

		prefixPrinter.Prefix.Text = "Total Issues:"
		prefixPrinter.Println(
			pterm.NewStyle(pterm.FgCyan, pterm.Bold).Sprint(category.TotalIssues),
		)
		fmt.Println(styles.Divider.Render("-----------------------------------"))
	}
}

func NewInsightsPerformanceView(apiResponse *types.ApiResponse) {
	for _, category := range apiResponse.Report.Categories {
		if category.Title == "Performance" {
			pterm.DefaultSection.WithStyle(sectionStyle).Println("Performance Report")

			tableData := pterm.TableData{
				{"Score", "Grade", "Total Issues"},
				{
					styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
					styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
					pterm.NewStyle(pterm.FgRed).Sprint(category.TotalIssues),
				},
			}

			pterm.DefaultTable.WithHeaderStyle(pterm.NewStyle(pterm.FgLightWhite)).WithHasHeader().WithBoxed().WithData(tableData).Render()

			pterm.DefaultSection.WithLevel(2).WithStyle(sectionStyle).Println("Performance Tests")

			testData := pterm.TableData{{"Test Title", "Status"}}
			for _, test := range category.Tests {
				row := []string{
					pterm.NewStyle(pterm.FgLightWhite, pterm.Bold).Sprint(test.Title),
					styles.RenderStyleForValue(test.Status),
				}
				testData = append(testData, row)
			}

			pterm.DefaultTable.WithData(testData).Render()
		}
	}
}

func NewInsightsDesignView(apiResponse *types.ApiResponse) {
	for _, category := range apiResponse.Report.Categories {
		if category.Title == "Design" {
			pterm.DefaultSection.WithStyle(sectionStyle).Println("Design Report")

			tableData := pterm.TableData{
				{"Score", "Grade", "Total Issues"},
				{
					styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
					styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
					pterm.NewStyle(pterm.FgRed).Sprint(category.TotalIssues),
				},
			}

			pterm.DefaultTable.WithHeaderStyle(pterm.NewStyle(pterm.FgLightWhite)).WithHasHeader().WithBoxed().WithData(tableData).Render()

			pterm.DefaultSection.WithLevel(2).WithStyle(sectionStyle).Println("Design Tests")
			testData := pterm.TableData{{"Test Title", "Status"}}
			for _, test := range category.Tests {
				row := []string{
					pterm.NewStyle(pterm.FgLightWhite, pterm.Bold).Sprint(test.Title),
					styles.RenderStyleForValue(test.Status),
				}
				testData = append(testData, row)
			}

			pterm.DefaultTable.WithData(testData).Render()
		}
	}
}

func NewInsightsSecurityView(apiResponse *types.ApiResponse) {
	for _, category := range apiResponse.Report.Categories {
		if category.Title == "Security" {
			pterm.DefaultSection.WithStyle(sectionStyle).Println("Security Report")
			tableData := pterm.TableData{
				{"Score", "Grade", "Total Issues"},
				{
					styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
					styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
					pterm.NewStyle(pterm.FgRed).Sprint(category.TotalIssues),
				},
			}

			pterm.DefaultTable.WithHeaderStyle(pterm.NewStyle(pterm.FgLightWhite)).WithHasHeader().WithBoxed().WithData(tableData).Render()

			testData := pterm.TableData{{"Test Title", "Status"}}
			for _, test := range category.Tests {
				row := []string{
					pterm.NewStyle(pterm.FgLightWhite, pterm.Bold).Sprint(test.Title),
					styles.RenderStyleForValue(test.Status),
				}
				testData = append(testData, row)
			}

			pterm.DefaultTable.WithData(testData).Render()
		}
	}
}

func NewInsightsFullView(apiResponse *types.ApiResponse) {
	NewInsightsDesignView(apiResponse)

	NewInsightsPerformanceView(apiResponse)

	NewInsightsSecurityView(apiResponse)
}
