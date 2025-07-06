package main

import (
	"cyoa/handler"
	"cyoa/parse"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file containing the CYOA story")
	flag.Parse()

	story, err := parse.ParseJson(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	for arc, chapter := range story {
		err = handler.Handler(chapter, arc, "./templates/index.html")
		if err != nil {
			log.Fatalln(err)
		}

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusFound)
	})

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Listening on %s", PORT)
}
