package views

import (
	"fmt"

	"github.com/treblle/treblle-cli/pkg/styles"
	"github.com/treblle/treblle-cli/pkg/types"
)

func NewApiInsightsView(apiResponse *types.ApiResponse) {

	fmt.Println(styles.Heading.Render("Details"))

	fmt.Println(
		styles.Title.Render("Title:  ") + styles.SubHeading.Render(apiResponse.Report.Title),
	)
	fmt.Println(
		styles.Title.Render("Description:  ") + styles.SubHeading.Render(fmt.Sprint(apiResponse.Report.Description)),
	)

	// Display basic report information with distinctive styles
	fmt.Println(
		styles.Title.Render("Share Url:  ") + styles.Link.Render(apiResponse.Report.ShareURL),
	)
	fmt.Println(
		styles.Title.Render("Base URL:  ") + styles.Link.Render(apiResponse.Report.BaseURL),
	)
	fmt.Println(
		styles.Title.Render("Endpoints:  ") + styles.SubHeading.Render(fmt.Sprint(apiResponse.Report.TotalEndpoints)),
	)
	fmt.Println(
		styles.Title.Render("Overall:  ") + styles.StyleForScore(apiResponse.Report.ScorePercentage).Render(fmt.Sprint(apiResponse.Report.ScorePercentage)),
	)

	fmt.Println(styles.Heading.Render("API Insights Report"))

	// // Category details
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

	fmt.Println(styles.Heading.Render("Technologies Discovered"))

	// Technology details
	for _, technology := range apiResponse.Report.Technologies {
		fmt.Println(
			styles.Title.Render("Name:  ") +
				styles.Value.Render(fmt.Sprintf("%v", technology.Name)),
		)

		fmt.Println(styles.Divider.Render("-----------------------------------"))
	}
}
