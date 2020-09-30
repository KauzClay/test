package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-commonmark/markdown"
)

func getFencedYAMLContent(tok markdown.Token) string {
	switch tok := tok.(type) {
	case *markdown.Fence:
		if tok.Params == "yaml" {
			return tok.Content
		}
	}
	return ""
}

// creates file if it doesn't exist
func writeToFile(name, content string) error {
	data := []byte(content)
	return ioutil.WriteFile(name, data, 0644)
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please specify input markdown file")
		fmt.Println("Usage: ./executable example-input.md")
		return
	}

	input := os.Args[1]
	filename := strings.TrimSuffix(input, filepath.Ext(input))
	output := filename

	doc, err := ioutil.ReadFile(input)

	if err != nil {
		panic(err)
	}

	//Parse the markdown
	md := markdown.New()
	tokens := md.Parse(doc)

	count := 0
	for _, t := range tokens {
		inlineCode := getFencedYAMLContent(t)

		if inlineCode != "" {
			outputFilename := output + "-" + strconv.Itoa(count) + ".yml"
			err := writeToFile(outputFilename, inlineCode)
			count += 1
			if err != nil {
				panic(err)
			}
		}
	}

}
