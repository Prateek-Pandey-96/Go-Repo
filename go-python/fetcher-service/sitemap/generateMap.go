package sitemap

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetLinks(urlParam string) []string {
	resp, err := http.Get(urlParam)
	if err != nil {
		log.Printf("err while fetching the given url %v", err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filterLinks(parseAndCleanLinks(resp.Body, base), base)
}

func parseAndCleanLinks(r io.Reader, base string) []string {
	links, _ := Parse(r)
	var hrefs []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link, "/"):
			hrefs = append(hrefs, base+link)
		case strings.HasPrefix(link, "http"):
			hrefs = append(hrefs, link)
		}
	}
	return hrefs
}

func filterLinks(links []string, base string) []string {
	ret := []string{}
	for _, link := range links {
		if strings.HasPrefix(link, base) {
			ret = append(ret, link)
		}
	}
	return ret
}
