package golinkparser

import (
	"io"

	"golang.org/x/net/html"
)

type link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := []link{}
	dfs(doc, &links)
	return links, nil
}

func dfs(node *html.Node, links *[]link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		var attrValue string
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				attrValue = attr.Val
				break
			}
		}
		*links = append(*links, link{Href: attrValue, Text: node.FirstChild.Data})
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, links)
	}
}
