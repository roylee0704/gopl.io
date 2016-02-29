package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/roylee0704/gopl.io/ch4/github"
)

// IssueServer keeps Github Issue tracker data.
type IssueServer struct {
	terms []string
	tmpl  *template.Template
}

// NewIssueServer returns new Server that serves github issue tracker query result.
func NewIssueServer(terms []string, tmpl *template.Template) *IssueServer {
	return &IssueServer{terms, tmpl}
}

func (i *IssueServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	issues, err := github.SearchIssues(i.terms)

	if err != nil {
		log.Print(err)
	}
	if err := i.tmpl.Execute(w, issues); err != nil {
		log.Print(err)
	}
}
