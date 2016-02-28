package github

import "time"

// IssuesURL is github api to search issues.
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is the data container for returned json
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
	Message    string
}

// Issue is the item container for IssuesSearchResult
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User has user info
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
