package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	XmlNs string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://github.com/prateek69", "the url you want to build the sitemap for")
	depth := flag.Int("depth", 2, "the depth you want to traverse to")
	flag.Parse()

	pages := BFS(*urlFlag, *depth)

	toXml := urlset{
		XmlNs: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}
