package main

import (
	"io"
	"net/http"
	"net/url"
	golinkparser "sitemapbuilder/parser"
	"strings"
)

func GetLinks(urlFlag string) []string {
	resp, err := http.Get(urlFlag)
	exit(err)
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return Filter(ParseAndCleanLinks(resp.Body, &base), withPrefix(&base))
}

func withPrefix(pfx *string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, *pfx)
	}
}

func ParseAndCleanLinks(r io.Reader, base *string) []string {
	links, _ := golinkparser.Parse(r)
	var hrefs []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			hrefs = append(hrefs, *base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			hrefs = append(hrefs, link.Href)
		}
	}
	return hrefs
}

func Filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func exit(err error) {
	if err != nil {
		panic(err)
	}
}
