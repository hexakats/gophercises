package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"urlshort"
)

func main() {
	yamlFlag := flag.String(
		"yaml",
		"",
		`Path to a yaml file (root is the package's directory) with the format:
		- path: /some-path
		  url: https://www.some-url.com/demo
		- path: /some-other-path
		  url: https://www.some-cooler-url.com/demo

		default: ""`,
	)

	jsonFlag := flag.String(
		"json",
		"",
		`Path to a json file (root is the package's directory) with the format:

		{
		  "path": "url",
		  "other-path": "other-url"
		}

		default: ""`,
	)

	boltDbFlag := flag.Bool(
		"bolt-db",
		false,
		`Set this flag to use a boltDB database instead of yaml/json file containing the directs`,
	)

	flag.Parse()

	if (*yamlFlag != "" && *jsonFlag != "" && *boltDbFlag) ||
		(*yamlFlag != "" && *jsonFlag != "") ||
		(*yamlFlag != "" && *boltDbFlag) ||
		(*jsonFlag != "" && *boltDbFlag) {
		fmt.Println("Error: Only one of --yaml, --json, or --bolt-db can be specified.")
		os.Exit(0)
	}

	if *yamlFlag == "" && *jsonFlag == "" && !*boltDbFlag {
		fmt.Println("Error: A flag must be specified.")
		os.Exit(0)
	}

	// Default mux
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	var handler http.HandlerFunc

	switch {
	case *yamlFlag != "":
		yaml, err := os.ReadFile(*yamlFlag)
		if err != nil {
			log.Println(err)
		}

		handler, err = urlshort.YAMLHandler([]byte(yaml), mapHandler)
		if err != nil {
			log.Println(err)
		}

	case *jsonFlag != "":
		json, err := os.ReadFile(*jsonFlag)
		if err != nil {
			log.Println(err)
		}

		handler, err = urlshort.JSONHandler([]byte(json), mapHandler)
		if err != nil {
			log.Println(err)
		}

	case *boltDbFlag:
		// TODO: Implement
		log.Fatalln("To be implemented.")
	default:
		break
	}

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	fmt.Println("Starting the server on :8080")

	err := http.ListenAndServe(":8080", handler)
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
