package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

func (p *Redirect) String() string {
	return fmt.Sprintf("%s,%s", p.Path, p.Url)
}

func (p *Redirect) Set(value string) error {
	// Split the input by a delimiter (e.g., comma)
	parts := strings.Split(value, " ")
	if len(parts) != 2 {
		return fmt.Errorf("expecting two values separated by a space")
	}
	p.Path = parts[0]
	p.Url = parts[1]
	return nil
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if destination, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, destination, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlFile []byte, fallback http.Handler) (http.HandlerFunc, error) {
	redirectArray := []Redirect{}
	redirectMap := make(map[string]string)

	err := yaml.Unmarshal(yamlFile, &redirectArray)
	if err != nil {
		log.Fatalln(err)
	}

	for _, redirect := range redirectArray {
		redirectMap[redirect.Path] = redirect.Url
	}

	return MapHandler(redirectMap, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
// ```json
// [
//
//	{
//	  "path": "/urlshort",
//	  "url": "https://github.com/gophercises/urlshort"
//	},
//	{
//	  "path": "/urlshort-solution",
//	  "url": "https://github.com/gophercises/urlshort/tree/solution"
//	},
//
// ]
//
// ```
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsonFile []byte, fallback http.Handler) (http.HandlerFunc, error) {
	redirectArray := []Redirect{}
	redirectMap := make(map[string]string)

	err := json.Unmarshal(jsonFile, &redirectArray)
	if err != nil {
		log.Fatalln(err)
	}

	for _, redirect := range redirectArray {
		redirectMap[redirect.Path] = redirect.Url
	}

	return MapHandler(redirectMap, fallback), nil
}

type Redirect struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url"  json:"url"`
}

// Struct name should be in lower case because its not being exported,
// This file only exports handler functions.
