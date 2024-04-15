package types

// ApiResponse comes back when a successful
// API Insights report has been created.
type ApiResponse struct {
	Message string `json:"message"`
	Report  Report `json:"report"`
}

type Industry struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// Report is the API
// Insights Report itself
type Report struct {
	UUID            string       `json:"uuid"`
	ShareURL        string       `json:"share_url"`
	Status          string       `json:"status"`
	Title           string       `json:"title"`
	BaseURL         string       `json:"base_url"`
	Description     string       `json:"description"`
	Version         string       `json:"version"`
	TotalEndpoints  int          `json:"total_endpoints"`
	CreatedAt       string       `json:"created_at"`
	ScorePercentage int          `json:"score_percentage"`
	AIReady         bool         `json:"is_ai_ready"`
	Grade           string       `json:"grade"`
	Industry        Industry     `json:"industry,omitempty"`
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
