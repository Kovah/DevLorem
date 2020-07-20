package main

import (
	"encoding/json"
	rice "github.com/GeertJohan/go.rice"
	"math/rand"
	"time"
)

type Source struct {
	Source     string   `json:"source"`
	Paragraphs []string `json:"paragraphs"`
}

func GetSources() []string {
	arr := make([]string, 27)

	arr[0] = "al-pacino.json"
	arr[1] = "anthony-hopkins.json"
	arr[2] = "charlize-theron.json"
	arr[3] = "christopher-walken.json"
	arr[4] = "clint-eastwood.json"
	arr[5] = "daisy-ridley.json"
	arr[6] = "george-washington.json"
	arr[7] = "harrison-ford.json"
	arr[8] = "hp-lovecraft.json"
	arr[9] = "jack-nicholson.json"
	arr[10] = "jennifer-lawrence.json"
	arr[11] = "jim-carrey.json"
	arr[12] = "jodie-foster.json"
	arr[13] = "johnny-depp.json"
	arr[14] = "leonardo-dicaprio.json"
	arr[15] = "master-yoda.json"
	arr[16] = "matt-damon.json"
	arr[17] = "mel-gibson.json"
	arr[18] = "michael-caine.json"
	arr[19] = "michel-houellebecq.json"
	arr[20] = "mikhail-gorbachev.json"
	arr[21] = "morgan-freeman.json"
	arr[22] = "president-obama.json"
	arr[23] = "robin-williams.json"
	arr[24] = "samuel-l-jackson.json"
	arr[25] = "tom-ellis.json"
	arr[26] = "tommy-lee-jones.json"

	return arr
}

func GetSourceContent(sourceFile string) (Source, error) {
	sourceBox := rice.MustFindBox("lorem")

	content, err := sourceBox.Bytes(sourceFile)
	Check(err)

	var source Source
	err = json.Unmarshal(content, &source)

	return source, err
}

func GetRandomContent(addParagraphs bool) Source {
	sources := GetSources()
	sourceFile := sources[rand.Intn(len(sources))]

	source, err := GetSourceContent(sourceFile)
	Check(err)

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

func GetNumLines(amount int, addParagraphs bool) Source {
	source := GetRandomContent(addParagraphs)

	// Fill results with random lines from the results by appending random lines
	for len(source.Paragraphs) < amount {
		rand.Seed(time.Now().UnixNano())
		source.Paragraphs = append(source.Paragraphs, source.Paragraphs[rand.Intn(len(source.Paragraphs))])
	}

	// Limit number of returned lines to the given amount
	source.Paragraphs = source.Paragraphs[0:amount]

	return source
}
