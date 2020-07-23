package main

import (
	"encoding/json"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	texttemplate "text/template"
)

var bindHost string

type Output struct {
	Source          Source
	ShowsParagraphs bool
}

func init() {
	rootCmd.AddCommand(httpCmd)

	httpCmd.Flags().StringVarP(&bindHost, "bind", "b", ":80", "Bind the HTTP server to a specific host and port, default is :80")
}

var httpCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the DevLorem website as a HTTP server.",
	Long:  `Run the DevLorem website as a HTTP server. The HTTP server also serves the web API.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Starting HTTP server for DevLorem on %v...\n", bindHost)

		if err := handleHttpServer(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func handleHttpServer() error {
	r := mux.NewRouter()

	// Handle static assets
	static := rice.MustFindBox("assets/dist")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(static.HTTPBox())))

	// Prepare both the HTML and plain text template
	templateBox := rice.MustFindBox("templates")
	templateString := templateBox.MustString("index.html")
	tmpl, err := template.New("index").Parse(templateString)
	Check(err)

	rawTmpl := texttemplate.Must(texttemplate.New("test").Parse("{{.Source}}\n{{range .Paragraphs}}{{.}}\n{{end}}"))

	// Handle the base webpage with generated paragraphs
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		source := GetRandomContent(false)

		output := Output{
			Source:          source,
			ShowsParagraphs: false,
		}

		err = tmpl.Execute(w, output)
		Check(err)
	})

	// Handle the base webpage with generated paragraphs and show the paragraph tags in the results
	r.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		source := GetRandomContent(true)

		output := Output{
			Source:          source,
			ShowsParagraphs: true,
		}

		err = tmpl.Execute(w, output)
		Check(err)
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
		Check(err)

		source := GetNumLines(amount, query.Get("paragraphs") != "true")

		if query.Get("format") == "text" {
			err = rawTmpl.Execute(w, source)
			Check(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(source)
			Check(err)
		}
	})

	// Additional API endpoint for returning proper errors if <amount> is larger than 99
	r.HandleFunc("/api/{amount:[0-9]{3}}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "You can only request up to 99 paragraphs in one request.")
		Check(err)
	})

	err = http.ListenAndServe(bindHost, r)
	return err
}
