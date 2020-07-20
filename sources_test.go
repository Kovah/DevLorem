package main

import (
	"fmt"
	"testing"
)

func TestSourceJsonFiles(t *testing.T) {
	sources := GetSources()

	fmt.Printf("Testing %v source files\n", len(sources))

	for _, sourceFile := range sources {
		source, err := GetSourceContent(sourceFile)

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
	source := GetNumLines(5, false)

	if len(source.Paragraphs) != 5 {
		t.Errorf("%v paragraphs returned, exptected %v", len(source.Paragraphs), 5)
	}

	source2 := GetNumLines(50, false)

	if len(source2.Paragraphs) != 50 {
		t.Errorf("%v paragraphs returned, exptected %v", len(source2.Paragraphs), 50)
	}
}
