package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int
	Items      []*Issue
}

type Issue struct {
	Number  int
	User    *User
	Title   string
	Created time.Time
	HTMLURL string
	State   string
}

type User struct {
	Login   string
	HTMLURL string
}

var issueList = template.Must(template.New("issuelist").Parse(`
 <h1>{{ .TotalCount}} issues</h1>
    <table>
        <tr style='text-align: left'>
            <th>#</th>
            <th>State</th>
            <th>User</th>
            <th>Title</th>
        </tr>
        {{range .Items}}
        <tr>
            <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
            <td>{{.State}}</td>
            <td><a href="{{.User.HTMLURL}}">{{.User.Login}}</a></td>
            <td><a href="{{.HTMLURL}}">{{.Title}}</a></td>
        </tr>
        {{end}}
    </table>
`))

func main() {
	http.HandleFunc("/issueList", func(writer http.ResponseWriter, request *http.Request) {
		result := IssuesSearchResult{
			TotalCount: 2,
			Items: []*Issue{
				{
					Number:  100,
					User:    &User{Login: "ljs", HTMLURL: "https://www.github.com"},
					Title:   "this is a personal github account",
					Created: time.Now(),
					HTMLURL: "https://www.baidu.com",
					State:   "starting",
				}, {
					Number:  101,
					User:    &User{Login: "jwt", HTMLURL: "https://linjianshu.github.io"},
					Title:   "this is a fuckful github account",
					Created: time.Now().Add(time.Hour * 24),
					HTMLURL: "https://www.tencent.com",
					State:   "ended",
				},
			},
		}
		issueList.Execute(writer, result)
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}
}
