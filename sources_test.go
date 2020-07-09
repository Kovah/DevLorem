package main

import (
	"fmt"
	"testing"
)

func TestSourceJsonFiles(t *testing.T) {
	sources := getSources()

	fmt.Printf("Testing %v source files\n", len(sources.Sources))

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

func TestNumLinesFunction(t *testing.T) {
	source := getNumLines(5, false)

	if len(source.Paragraphs) != 5 {
		t.Errorf("%v paragraphs returned, exptected %v", len(source.Paragraphs), 5)
	}

	source2 := getNumLines(50, false)

	if len(source2.Paragraphs) != 50 {
		t.Errorf("%v paragraphs returned, exptected %v", len(source2.Paragraphs), 50)
	}
}
