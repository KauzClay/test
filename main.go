package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-commonmark/markdown"
)

//snippet represents the snippet we will output.
type snippet struct {
	content string
	lang    string
}

//getSnippet extract only code snippet from markdown object.
func getSnippet(tok markdown.Token) snippet {
	switch tok := tok.(type) {
	case *markdown.CodeBlock:
		return snippet{
			tok.Content,
			"code",
		}
	case *markdown.CodeInline:
		return snippet{
			tok.Content,
			"code inline",
		}
	case *markdown.Fence:
		return snippet{
			tok.Content,
			tok.Params,
		}
	}
	return snippet{}
}

//readFromWeb call the given url and return the content of the readme.
func readFromWeb(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {

	doc, err := ioutil.ReadFile("test.md")

	if err != nil {
		panic(err)
	}

	//Parse the markdown
	md := markdown.New()
	tokens := md.Parse(doc)

	fmt.Println(len(tokens))
	//Print the result
	for _, t := range tokens {
		snippet := getSnippet(t)

		fmt.Println(snippet)
	}

}
