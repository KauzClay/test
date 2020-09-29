package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang-commonmark/markdown"
)

func getFencedYAMLContent(tok markdown.Token) string {
	switch tok := tok.(type) {
	case *markdown.Fence:
		fmt.Println(tok)
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

	input := os.Args[1]
	output := flag.String("output-filename", "output.yml", "If not specified, will output to output.yml")

	doc, err := ioutil.ReadFile(input)

	if err != nil {
		panic(err)
	}

	//Parse the markdown
	md := markdown.New()
	tokens := md.Parse(doc)

	for _, t := range tokens {
		inlineCode := getFencedYAMLContent(t)

		if inlineCode != "" {
			err := writeToFile(*output, inlineCode)
			if err != nil {
				panic(err)
			}
		}
	}

}
