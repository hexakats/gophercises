package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"urlshort/handler"
)

func main() {
	yamlFlag := flag.String(
		"yaml",
		"",
		`Path to a yaml file with the format:
		- path: /some-path
		  url: https://www.some-url.com/demo
		- path: /some-other-path
		  url: https://www.some-cooler-url.com/demo

		default: ""`,
	)

	jsonFlag := flag.String(
		"json",
		"",
		`Path to a json file with the format:

		{
		  "path": "url",
		  "other-path": "other-url"
		}

		default: ""`,
	)

	flag.Parse()

	if *yamlFlag != "" && *jsonFlag != "" {
		fmt.Println("Error: Only one of --yaml, --json, or --bolt-db can be specified.")
		os.Exit(0)
	}

	if *yamlFlag == "" && *jsonFlag == "" {
		fmt.Println("Error: A flag must be specified.")
		os.Exit(0)
	}

	// Default mux
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{}
	mapHandler := handler.MapHandler(pathsToUrls, mux)
	var httpHandler http.HandlerFunc

	switch {
	case *yamlFlag != "":
		yaml, err := os.ReadFile(*yamlFlag)
		if err != nil {
			log.Println(err)
		}

		httpHandler, err = handler.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			log.Println(err)
		}

	case *jsonFlag != "":
		json, err := os.ReadFile(*jsonFlag)
		if err != nil {
			log.Println(err)
		}

		httpHandler, err = handler.JSONHandler([]byte(json), mapHandler)
		if err != nil {
			log.Println(err)
		}
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	fmt.Println("Starting the server on :8080")

	err := http.ListenAndServe(":8080", httpHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
