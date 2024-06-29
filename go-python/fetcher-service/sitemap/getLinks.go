package sitemap

import (
	"io"

	"golang.org/x/net/html"
)

func Parse(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	hrefs := []string{}
	dfs(doc, &hrefs)
	return hrefs, nil
}

// Perform DFS and find all the a tags on the page
func dfs(node *html.Node, hrefs *[]string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		var attrValue string
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				attrValue = attr.Val
				break
			}
		}
		*hrefs = append(*hrefs, attrValue)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, hrefs)
	}
}
