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

	ElementByID(doc, id)
	if targetNode == nil {
		fmt.Println("Not found")
	} else {
		fmt.Printf("'%s' element found\n", targetNode.Data)
	}
}

var targetID string
var targetNode *html.Node

// ElementByID finds the first HTML element with the specified id attribute
func ElementByID(doc *html.Node, id string) {
	targetID = id
	targetNode = nil
	forEachNode(doc, startElement, nil)
}

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == targetID {
				targetNode = n
				return false
			}
		}
	}
	return true
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil && !pre(n) {
		return false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isContinue := forEachNode(c, pre, post); !isContinue {
			return false
		}
	}

	if post != nil && !post(n) {
		return false
	}
	return true
}
