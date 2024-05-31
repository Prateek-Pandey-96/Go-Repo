package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prateek69/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlFile := flag.String("yaml", "paths.yaml", "enter a yaml file")
	jsonFile := flag.String("json", "paths.json", "enter a json file")
	flag.Parse()

	yaml, err := os.ReadFile(*yamlFile)
	exit(err)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	exit(err)

	// Build the JsonHandler using yamlHandler as the fallback
	json, err := os.ReadFile(*jsonFile)
	exit(err)

	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	exit(err)

	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func exit(err error) {
	if err != nil {
		panic(err)
	}
}
