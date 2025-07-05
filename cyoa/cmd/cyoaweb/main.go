package main

import (
	"cyoa/parse"
	"cyoa/story"
	"flag"
	"html/template"
	"log"
	"net/http"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file containing the CYOA story")
	flag.Parse()

	story, err := parse.ParseJson(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

	chapter := "intro"

	tmpl := template.Must(template.ParseFiles("./template/story.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := StoryPage{
			Title:   story[chapter].Title,
			Story:   story[chapter].Story,
			Options: story[chapter].Options,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Fatalln(err)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

type StoryPage struct {
	Title   string
	Story   []string
	Options []story.Option
}
