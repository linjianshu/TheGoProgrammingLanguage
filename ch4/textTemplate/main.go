package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

const tmpl = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------
Number : {{.Number}}
User: {{.User.Login}}
Title : {{.Title |printf "%.64s"}}
Age : {{.Created | daysAgo}} days
{{end}}
`

type IssuesSearchResult struct {
	TotalCount int
	Items      []*Issue
}

type Issue struct {
	Number  int
	User    *User
	Title   string
	Created time.Time
}

type User struct {
	Login string
}

func main() {
	var daysAgo func(now time.Time) int
	daysAgo = func(now time.Time) int {
		return int(time.Since(now).Hours() / 24)
	}

	template := template.Must(template.New("tmpl").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(tmpl))
	result := IssuesSearchResult{
		TotalCount: 2,
		Items: []*Issue{
			{
				Number:  100,
				User:    &User{Login: "ljs"},
				Title:   "this is a personal github account",
				Created: time.Now(),
			}, {

				Number:  101,
				User:    &User{Login: "jwt"},
				Title:   "this is a fuckful github account",
				Created: time.Now().Add(time.Hour * 24),
			},
		},
	}
	err := template.Execute(os.Stdout, result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
