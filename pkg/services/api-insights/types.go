package apiinsights

import (
	"time"
)

type Report struct {
	UUID            string       `json:"uuid"`
	ShareURL        string       `json:"share_url"`
	Status          string       `json:"status"`
	Title           string       `json:"title"`
	BaseURL         string       `json:"base_url"`
	Description     string       `json:"description"`
	Version         string       `json:"version"`
	TotalEndpoints  int          `json:"total_endpoints"`
	CreatedAt       time.Time    `json:"created_at"`
	ScorePercentage int          `json:"score_percentage"`
	Grade           string       `json:"grade"`
	Industry        string       `json:"industry"`
	Categories      []Category   `json:"categories"`
	Technologies    []Technology `json:"technologies"`
	DeleteToken     string       `json:"delete_token"`
}

type Category struct {
	Title           string `json:"title"`
	ScorePercentage int    `json:"score_percentage"`
	Grade           string `json:"grade"`
	TotalIssues     int    `json:"total_issues"`
	Tests           []Test `json:"tests"`
}

type Test struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Technology struct {
	Name     string `json:"name"`
	IconPath string `json:"icon_path"`
}

type APIResponse struct {
	Message string `json:"message"`
	Report  Report `json:"report"`
}
