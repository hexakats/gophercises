package parse

import (
	"cyoa/story"
	"encoding/json"
	"io"
	"os"
)

func ParseJson(filename *string) (story.Story, error) {
	file, err := os.Open(*filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var story story.Story

	err = json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}

	return story, err
}
