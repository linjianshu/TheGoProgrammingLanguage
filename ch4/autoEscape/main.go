package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	const tmpl = `<p> A:{{.A}}</p> {{.C}} <p>B: {{.B}}</p>`
	t := template.Must(template.New("template").Parse(tmpl))
	var data struct {
		A string
		//template.HTML 将会自动转义 没有使用这种类型的不会自动转义 防止脚本注入引发安全问题
		B template.HTML
		C template.HTML
	}
	d := data
	d.A = "<b>Hello !</b>"
	d.B = "<b>Hello !</b>"
	d.C = "</br>"

	http.HandleFunc("/getFuck", func(writer http.ResponseWriter, request *http.Request) {
		t.Execute(writer, d)
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}

}
