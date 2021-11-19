package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Reader implements the io.Reader
type Reader struct {
	doc []byte
}

// Read implements the interface of the io.Reader
func (r *Reader) Read(p []byte) (n int, err error) {
	d := []byte(r.doc)
	if len(d) == 0 {
		err = io.EOF
		return
	}
	end := len(p)
	if len(d) < len(p) {
		end = len(d)
	}
	n = copy(p, d[:end])
	r.doc = d[end:]
	return
}

// NewReader returns a new Reader
func NewReader(s string) *Reader {
	return &Reader{[]byte(s)}
}

const doc = `
<!DOCTYPE html>
<html>
  <head>
    <title>My First HTML</title>
    <meta charset="UTF-8">
  </head>
  <body>
    <h1>My First Heading</h1>
    <p>My first paragraph.</p>
    <h2>An Unordered HTML List</h2>
    <ul>
      <li>Coffee</li>
      <li>Tea</li>
      <li>Milk</li>
    </ul>
    <h2>An Ordered HTML List</h2>
    <ol>
      <li>Coffee</li>
      <li>Tea</li>
      <li>Milk</li>
    </ol>
  </body>
  </html>
`

func main() {
	r := NewReader(doc)

	outline(r)
}

func outline(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
