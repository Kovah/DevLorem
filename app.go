package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	texttemplate "text/template"
	"time"
)

type Sources struct {
	Sources []string
}

type Source struct {
	Source     string
	Paragraphs []string
}

type Output struct {
	Source          Source
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

func getRandomContent(addParagraphs bool) Source {
	sources := getSources()
	sourceFile := sources.Sources[rand.Intn(len(sources.Sources))]

	content, err := ioutil.ReadFile("lorem/" + sourceFile)
	check(err)

	var source Source
	json.Unmarshal(content, &source)

	if addParagraphs {
		for i, paragraph := range source.Paragraphs {
			source.Paragraphs[i] = "<p>" + paragraph + "</p>"
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(source.Paragraphs), func(i, j int) {
		source.Paragraphs[i], source.Paragraphs[j] = source.Paragraphs[j], source.Paragraphs[i]
	})

	return source
}

func getNumLines(amount int, stripParagraphs bool) Source {
	source := getRandomContent(stripParagraphs)

	// Fill results with random lines from the results by appending random lines
	for len(source.Paragraphs) < amount {
		rand.Seed(time.Now().UnixNano())
		source.Paragraphs = append(source.Paragraphs, source.Paragraphs[rand.Intn(len(source.Paragraphs))])
	}

	// Limit number of returned lines to the given amount
	source.Paragraphs = source.Paragraphs[0:amount]

	return source
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
		source := getRandomContent(false)

		output := Output{
			Source:          source,
			ShowsParagraphs: false,
		}

		tmpl.Execute(w, output)
	})

	// Handle the base webpage with generated paragraphs and show the paragraph tags in the results
	r.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		source := getRandomContent(true)

		output := Output{
			Source:          source,
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

		source := getNumLines(amount, query.Get("paragraphs") != "true")

		if query.Get("format") == "text" {
			rawTmpl.Execute(w, source)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(source)
		}
	})

	// Additional API endpoint for returning proper errors if <amount> is larger than 99
	r.HandleFunc("/api/{amount:[0-9]{3}}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "You can only request up to 99 paragraphs in one request.")
	})

	err := http.ListenAndServe(":80", r)
	check(err)
}
