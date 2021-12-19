package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type link struct {
	url   string
	depth int
}

var tokens = make(chan struct{}, 20)

func toLinks(list []string, depth int) []link {
	links := []link{}
	for _, l := range list {
		links = append(links, link{l, depth})
	}
	return links
}

func crawl(url string, depth int) []link {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return toLinks(list, depth+1)
}

func main() {
	var maxDepth int
	flag.IntVar(&maxDepth, "depth", 1, "Max depth for clawling.")
	flag.Parse()

	worklist := make(chan []link)
	var n int // number of pending sends to worklist
	var depth int

	// Start with the command-line arguments.
	n++
	go func() { worklist <- toLinks(flag.Args(), depth) }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		fmt.Println(n)
		links := <-worklist
		for _, l := range links {
			if !seen[l.url] && l.depth < maxDepth {
				seen[l.url] = true
				n++
				go func(l link) {
					worklist <- crawl(l.url, l.depth)
				}(l)
			}
		}
	}
}
