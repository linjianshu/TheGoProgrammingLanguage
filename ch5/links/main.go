package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	s, err := Extract("https://linjianshu.github.io/")
	if err != nil {
		fmt.Println(err)
	}

	for i, url := range s {
		fmt.Printf("%v :\t %s\n", i, url)
	}
}
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s : %s ", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML : %v", url, err)
	}

	var links []string
	var visitNode func(node *html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(doc)

	return links, nil

}
