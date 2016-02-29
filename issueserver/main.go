package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
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

func main() {

	http.Handle("/", NewIssueServer(os.Args[1:], html))
	log.Fatal(http.ListenAndServe("localhost:8989", nil))

}
