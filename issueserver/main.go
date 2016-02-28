package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/roylee0704/gopl.io/ch4/github"
)

const templ = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
  <h1>{{.TotalCount}} issues</h1>
  <table>
    <tr>
      <th>#</th>
      <th>State</th>
      <th>User</th>
      <th>Title</th>
    </tr>
    {{range .Items}}
    <tr>
      <td>{{.Number}}</td>
      <td>{{.State}}</td>
      <td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
      <td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
    </tr>
    {{end}}
  </table>
</body>
</html>`

var html = template.Must(template.New("issuelist").Parse(templ))

func ServeSomething(w http.ResponseWriter, r *http.Request) {
	issues, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := html.Execute(w, issues); err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", ServeSomething)
	if err := http.ListenAndServe("localhost:8989", nil); err != nil {
		log.Fatal(err)
	}
}
