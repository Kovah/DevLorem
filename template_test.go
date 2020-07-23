package main

import (
	"bytes"
	rice "github.com/GeertJohan/go.rice"
	"html/template"
	"strings"
	"testing"
)

func TestTemplateParsing(t *testing.T) {
	templateBox := rice.MustFindBox("templates")
	templateString := templateBox.MustString("index.html")
	tmpl, err := template.New("index").Parse(templateString)

	if err != nil {
		t.Error(err)
	}

	output := Output{
		Source: Source{
			Source:     "Example Source",
			Paragraphs: []string{"Test Paragraph", "Test Paragraph 2"},
		},
		ShowsParagraphs: false,
	}

	// Using a byte buffer to capture the output of the template and save it to a string afterwards
	var tmplBuffer bytes.Buffer
	err = tmpl.Execute(&tmplBuffer, output)
	html := tmplBuffer.String()

	if err != nil {
		t.Error(err)
	}

	if len(html) == 0 {
		t.Log("Output of the template is empty, should contain HTML output")
		t.FailNow()
	}

	if !strings.Contains(html, "Example Source") {
		t.Error("Source is missing in the template output")
	}

	if !strings.Contains(html, "Test Paragraph") {
		t.Error("Paragraph is missing in the template output")
	}
}
