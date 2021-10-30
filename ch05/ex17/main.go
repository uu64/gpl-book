package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func contains(slice []string, val string) bool {
	for _, v := range slice {
		if strings.Compare(v, val) == 0 {
			return true
		}
	}

	return false
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	var search func(n *html.Node, name ...string)

	search = func(n *html.Node, name ...string) {
		if n.Type == html.ElementNode {
			if contains(name, n.Data) {
				elements = append(elements, n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			search(c, name...)
		}
	}

	search(doc, name...)
	return elements
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		images := ElementsByTagName(doc, "img")
		fmt.Printf("%d elements was found\n", len(images))
	}
}
