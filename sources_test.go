package main

import (
	"testing"
)

func TestSourceJsonFiles(t *testing.T) {
	sources := getSources()

	for _, sourceFile := range sources.Sources {
		source, err := getSourceContent(sourceFile)

		if err != nil {
			t.Logf("Source file %v contains invalid Json", sourceFile)
			t.FailNow()
		}

		if source.Source == "" {
			t.Errorf("Source field inside the %v file is empty", sourceFile)
		}

		if len(source.Paragraphs) == 0 {
			t.Errorf("Source file %v contains no paragraphs", sourceFile)
		}

		for index, paragraph := range source.Paragraphs {
			if paragraph == "" {
				t.Errorf("Paragraph %v the %v file is empty", index, sourceFile)
			}
		}
	}
}
