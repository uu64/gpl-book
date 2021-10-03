package issue

import "time"

const baseURL = "https://api.github.com"

// Issue is github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
	Body      string    // in Markdown format
}

// User is github user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
