package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	id := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to fetch %s: %v", url, err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("failed to parse document: %v", err)
	}

	n := ElementByID(doc, id)
	fmt.Println(n)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	_, n := forEachNode(doc, nil, nil)
	return n
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (bool, *html.Node) {
	if pre != nil && !pre(n) {
		return false, n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isContinue, lastNode := forEachNode(c, pre, post); !isContinue {
			return false, lastNode
		}
	}

	if post != nil && !post(n) {
		return false, n
	}
	return true, nil
}
