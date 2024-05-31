package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. parse the yaml
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	// 2. invert yaml array into map
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}
	// 3. return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsAndUrls pathAndUrl
	err := json.Unmarshal(jsonBytes, &pathsAndUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)

	for idx, path := range pathsAndUrls.Paths {
		pathsToUrls[path] = pathsAndUrls.Urls[idx]
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type pathAndUrl struct {
	Paths []string `json:"paths"`
	Urls  []string `json:"urls"`
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
