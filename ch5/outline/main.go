package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	buffer := bytes.Buffer{}
	//不能含有 /n 转义字符
	buffer.WriteString(`<div class="container-fluid"><div class="navbar-header"><button aria-controls="navbar" aria-expanded="false" class="navbar-toggle collapsed" data-target="#navbar" data-toggle="collapse" type="button"><span class="sr-only">Toggle navigation</span><span class="icon-bar"></span><span class="icon-bar"></span><span class="icon-bar"></span></button><a class="navbar-brand" href="#">安全件解析规则清单</a></div><div class="navbar-collapse collapse" id="navbar"><ul class="nav navbar-nav navbar-right"><li><a href="/CheckPart">主面板</a></li><li><a href="https://linjianshu.github.io" target="_blank">帮助</a></li></ul><form action="/Search" class="navbar-form navbar-right"><input class="form-control" name="keyword" placeholder="Search..." type="text"></form></div></div>`)
	parse, err := html.Parse(&buffer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline : %v\n", err)
		os.Exit(1)
	}
	outline(nil, parse)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
