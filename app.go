package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	texttemplate "text/template"
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

type JsonOutput struct {
	Source     string   `json:"source"`
	Paragraphs []string `json:"paragraphs"`
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

func getNumLines(amount int, stripParagraphs bool) (string, []string) {
	source, lines := getRandomContent(stripParagraphs)

	// Fill results with random lines from the results by appending random lines
	for len(lines) < amount {
		rand.Seed(time.Now().UnixNano())
		lines = append(lines, lines[rand.Intn(len(lines))])
	}

	// Limit number of returned lines to the given amount
	lines = lines[0:amount]

	return source, lines
}

func main() {

	r := mux.NewRouter()

	// Handle static assets
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/dist"))))

	// Prepare both the HTML and plain text template
	tmpl := template.Must(template.ParseFiles("template.html"))
	rawTmpl := texttemplate.Must(texttemplate.New("test").Parse("{{.Source}}\n{{range .Paragraphs}}{{.}}\n{{end}}"))

	// Handle the base webpage with generated paragraphs
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		source, content := getRandomContent(true)

		output := Output{
			Source:          strings.TrimSuffix(source, ".txt"),
			Paragraphs:      content,
			ShowsParagraphs: false,
		}

		tmpl.Execute(w, output)
	})

	// Handle the base webpage with generated paragraphs and show the paragraph tags in the results
	r.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		source, content := getRandomContent(false)

		output := Output{
			Source:          strings.TrimSuffix(source, ".txt"),
			Paragraphs:      content,
			ShowsParagraphs: true,
		}

		tmpl.Execute(w, output)
	})

	// Handle API calls
	// /api/<amount>[?paragraphs=true][&format=text]
	// <amount> specifies how many paragraphs to return, maximum is 99
	// paragraphs=true tells the API to include paragraph tags in the results
	// format=text tells the API to return plain text instead of Json
	r.HandleFunc("/api/{amount:[0-9]{1,2}}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		query := r.URL.Query()

		amount, err := strconv.Atoi(vars["amount"])
		check(err)

		source, content := getNumLines(amount, query.Get("paragraphs") != "true")

		output := JsonOutput{
			Source:     strings.TrimSuffix(source, ".txt"),
			Paragraphs: content,
		}

		if query.Get("format") == "text" {
			rawTmpl.Execute(w, output)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(output)
		}
	})

	err := http.ListenAndServe(":80", r)
	check(err)
}
