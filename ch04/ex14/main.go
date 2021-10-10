package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/uu64/gpl-book/ch04/ex14/github"
)

const addr = "localhost:8000"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Milestone</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  {{ if .Milestone}}
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Title}}</a></td>
  {{ else }}
  <td>No milestone</td>
  {{ end }}
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

type searchParams struct {
	owner string
	repo  string
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		owner, repo, err := parseURL(r.URL)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		is := parseQuery(r.URL)

		query := []string{fmt.Sprintf("repo:%s/%s", owner, repo), fmt.Sprintf("is:%s", is)}
		result, err := github.SearchIssues(query)
		if err != nil {
			log.Print(err)
		}

		showIssueList(w, result)
	}
	http.HandleFunc("/issues/", handler)
	fmt.Printf("listen and serve... http://%s\n", addr)
	fmt.Printf("path: /issues/<owner>/<repo> :  Show issues in the github.com/<owner>/<repo>\n")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func parseURL(url *url.URL) (string, string, error) {
	var owner, repo string
	items := strings.Split(url.Path, "/")
	// path format: /isssues/owner/repo
	if len(items) < 4 {
		return owner, repo, fmt.Errorf("error: please set the owner and repo in url path")
	}

	owner = items[2]
	repo = items[3]
	if owner == "" || repo == "" {
		return owner, repo, fmt.Errorf("error: please set the owner and repo in url path")
	}

	return owner, repo, nil
}

func parseQuery(url *url.URL) string {
	var is string
	q := url.Query()
	is = q.Get("is")
	return is
}

func showIssueList(out io.Writer, result *github.IssuesSearchResult) {
	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}
