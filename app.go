package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
	"strconv"
	texttemplate "text/template"
)

type Output struct {
	Source          Source
	ShowsParagraphs bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
