package github

import "time"

const issuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is an object that represents the search result
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is an object that represents a github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	Milestone *Milestone
}

// Milestone is an object that represents a github milestone
type Milestone struct {
	HTMLURL     string `json:"html_url"`
	Title       string
	Description string
}

// User is an object that represents a github user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
