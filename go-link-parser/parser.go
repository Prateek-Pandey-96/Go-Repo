package golinkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := []Link{}
	dfs(doc, &links)
	return links, nil
}

func dfs(node *html.Node, links *[]Link) {
	if node.Type == html.ElementNode && node.Data == "a" {
		var attrValue string
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				attrValue = attr.Val
				break
			}
		}
		*links = append(*links, Link{Href: attrValue, Text: strings.TrimSpace(node.FirstChild.Data)})
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, links)
	}
}
