package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Sources struct {
	Sources []string
}

type Source struct {
	Source     string
	Paragraphs []string
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

func getSourceContent(sourceFile string) (Source, error) {
	content, err := ioutil.ReadFile("lorem/" + sourceFile)
	check(err)

	var source Source
	err = json.Unmarshal(content, &source)

	return source, err
}

func getRandomContent(addParagraphs bool) Source {
	sources := getSources()
	sourceFile := sources.Sources[rand.Intn(len(sources.Sources))]

	source, err := getSourceContent(sourceFile)
	check(err)

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

func getNumLines(amount int, addParagraphs bool) Source {
	source := getRandomContent(addParagraphs)

	// Fill results with random lines from the results by appending random lines
	for len(source.Paragraphs) < amount {
		rand.Seed(time.Now().UnixNano())
		source.Paragraphs = append(source.Paragraphs, source.Paragraphs[rand.Intn(len(source.Paragraphs))])
	}

	// Limit number of returned lines to the given amount
	source.Paragraphs = source.Paragraphs[0:amount]

	return source
}
