package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

type Sources struct {
	Sources []string
}

type Output struct {
	Source          string
	Paragraphs      string
	ShowsParagraphs bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSources() Sources {
	sourceDir, err := ioutil.ReadDir("lorem")
	check(err)

	sources := Sources{}
	for _, entry := range sourceDir {
		sources.Sources = append(sources.Sources, entry.Name())
	}

	return sources
}

func getRandomContent() (string, string) {
	sources := getSources()
	source := sources.Sources[rand.Intn(len(sources.Sources))]

	content, err := ioutil.ReadFile("lorem/" + source)
	check(err)

	parsedContent := string(content)

	return source, parsedContent
}

func main() {

	tmpl := template.Must(template.ParseFiles("template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		source, content := getRandomContent()

		content = strings.ReplaceAll(content, "<p>", "")
		content = strings.ReplaceAll(content, "</p>", "")

		output := Output{
			Source:          strings.TrimSuffix(source, ".txt"),
			Paragraphs:      content,
			ShowsParagraphs: false,
		}

		tmpl.Execute(w, output)
	})

	http.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		source, content := getRandomContent()

		output := Output{
			Source:          strings.TrimSuffix(source, ".txt"),
			Paragraphs:      content,
			ShowsParagraphs: true,
		}

		tmpl.Execute(w, output)
	})

	http.ListenAndServe(":80", nil)
}
