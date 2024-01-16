package views

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/treblle/treblle-cli/pkg/styles"
	"github.com/treblle/treblle-cli/pkg/types"
)

func ShowInsightsDetails(apiResponse *types.ApiResponse) {
	pterm.DefaultSection.Println("Details")

	printer := pterm.PrefixPrinter{
		MessageStyle: &pterm.ThemeDefault.InfoMessageStyle,
		Prefix: pterm.Prefix{
			Style: &pterm.ThemeDefault.InfoPrefixStyle,
			Text:  "INFO",
		},
	}

	printer.Prefix.Text = "Title"
	printer.Println(
		styles.SubHeading.Render(apiResponse.Report.Title),
	)

	printer.Prefix.Text = "Description"
	printer.Println(
		styles.SubHeading.Render(fmt.Sprint(apiResponse.Report.Description)),
	)

	printer.Prefix.Text = "View in Browser"
	printer.Println(
		styles.Link.Render(apiResponse.Report.ShareURL),
	)

	printer.Prefix.Text = "Base URL"
	printer.Println(
		styles.Link.Render(apiResponse.Report.BaseURL),
	)

	printer.Prefix.Text = "Endpoints"
	printer.Println(
		styles.SubHeading.Render(fmt.Sprint(apiResponse.Report.TotalEndpoints)),
	)

	printer.Prefix.Text = "Overall"
	printer.Println(
		styles.StyleForScore(apiResponse.Report.ScorePercentage).Render(fmt.Sprint(apiResponse.Report.ScorePercentage)),
	)
}

func ShowInsightsTechnologyDiscovery(apiResponse *types.ApiResponse) {
	pterm.DefaultSection.WithLevel(2).Println("Technologies Discovered")

	testData := pterm.TableData{{"Technology Name"}}
	for _, technology := range apiResponse.Report.Technologies {
		row := []string{
			styles.Title.Render(technology.Name),
		}
		testData = append(testData, row)
	}

	pterm.DefaultTable.WithData(testData).Render()

	fmt.Println(styles.Divider.Render("-----------------------------------"))
}

func NewApiInsightsView(apiResponse *types.ApiResponse) {
	pterm.DefaultSection.Println("API Insights Report")

	for _, category := range apiResponse.Report.Categories {
		fmt.Println(
			styles.Title.Render("Category:  ") +
				styles.Value.Render(fmt.Sprintf("%v", category.Title)),
		)
		fmt.Println(
			styles.Title.Render("Score:  ") +
				styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
		)
		fmt.Println(
			styles.Title.Render("Grade:  ") +
				styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
		)
		fmt.Println(
			styles.Title.Render("Total Issues:  ") +
				styles.Value.Render(fmt.Sprintf("%v", category.TotalIssues)),
		)
		fmt.Println(styles.Divider.Render("-----------------------------------"))
	}
}

func NewInsightsPerformanceView(apiResponse *types.ApiResponse) {
	for _, category := range apiResponse.Report.Categories {
		if category.Title == "Performance" {
			pterm.DefaultSection.Println("Performance Report")

			tableData := pterm.TableData{
				{"Score", "Grade", "Total Issues"},
				{
					styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
					styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
					styles.Value.Render(fmt.Sprintf("%v", category.TotalIssues)),
				},
			}

			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()

			pterm.DefaultSection.WithLevel(2).Println("Performance Tests")

			testData := pterm.TableData{{"Test Title", "Status"}}
			for _, test := range category.Tests {
				row := []string{
					styles.Title.Render(test.Title),
					styles.Value.Render(test.Status),
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
			pterm.DefaultSection.WithLevel(2).Println("Design Report")

			tableData := pterm.TableData{
				{"Score", "Grade", "Total Issues"},
				{
					styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
					styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
					styles.Value.Render(fmt.Sprintf("%v", category.TotalIssues)),
				},
			}

			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()

			pterm.DefaultSection.WithLevel(2).Println("Design Tests")

			testData := pterm.TableData{{"Test Title", "Status"}}
			for _, test := range category.Tests {
				row := []string{
					styles.Title.Render(test.Title),
					styles.Value.Render(test.Status),
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
			pterm.DefaultSection.WithLevel(2).Println("Security Report")
			tableData := pterm.TableData{
				{"Score", "Grade", "Total Issues"},
				{
					styles.StyleForScore(category.ScorePercentage).Render(fmt.Sprintf("%v", category.ScorePercentage)),
					styles.StyleForGrade(category.Grade).Render(fmt.Sprintf("%v", category.Grade)),
					styles.Value.Render(fmt.Sprintf("%v", category.TotalIssues)),
				},
			}

			pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()

			pterm.DefaultSection.WithLevel(2).Println("Security Tests")

			testData := pterm.TableData{{"Test Title", "Status"}}
			for _, test := range category.Tests {
				row := []string{
					styles.Title.Render(test.Title),
					styles.Value.Render(test.Status),
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
