package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Sources struct {
	Sources []string
}

type Output struct {
	Source          string
	Paragraphs      []string
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

func getRandomContent(stripParagraphs bool) (string, []string) {
	sources := getSources()
	source := sources.Sources[rand.Intn(len(sources.Sources))]

	content, err := ioutil.ReadFile("lorem/" + source)
	check(err)

	parsedContent := string(content)

	if stripParagraphs {
		parsedContent = strings.ReplaceAll(parsedContent, "<p>", "")
		parsedContent = strings.ReplaceAll(parsedContent, "</p>", "")
	}

	parsedLines := strings.Split(parsedContent, "\n")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(parsedLines), func(i, j int) {
		parsedLines[i], parsedLines[j] = parsedLines[j], parsedLines[i]
	})

	return source, parsedLines
}

func main() {

	tmpl := template.Must(template.ParseFiles("template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		source, content := getRandomContent(true)

		output := Output{
			Source:          strings.TrimSuffix(source, ".txt"),
			Paragraphs:      content,
			ShowsParagraphs: false,
		}

		tmpl.Execute(w, output)
	})

	http.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		source, content := getRandomContent(false)

		output := Output{
			Source:          strings.TrimSuffix(source, ".txt"),
			Paragraphs:      content,
			ShowsParagraphs: true,
		}

		tmpl.Execute(w, output)
	})

	http.ListenAndServe(":80", nil)
}
