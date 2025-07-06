package handler

import (
	"cyoa/story"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Handler(chapter story.Chapter, arc string, filePath string) error {
	tmpl := template.Must(template.ParseFiles(filePath))

	http.HandleFunc(fmt.Sprintf("/%s", arc), func(w http.ResponseWriter, r *http.Request) {
		data := StoryPage{
			Title:   chapter.Title,
			Story:   chapter.Story,
			Options: chapter.Options,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Fatalln(err)
		}
	})
	return nil
}

type StoryPage struct {
	Title   string
	Story   []string
	Options []story.Option
}
