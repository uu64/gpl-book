package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		content := strings.TrimSpace(n.Data)
		if len(content) != 0 {
			input := bufio.NewScanner(strings.NewReader(content))
			input.Split(bufio.ScanWords)
			for input.Scan() {
				words++
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}

func main() {
	words, images, err := CountWordsAndImages("https://golang.org")
	if err != nil {
		log.Fatal("error: %w", err)
	}
	fmt.Printf("words: %d, images: %d\n", words, images)
}
