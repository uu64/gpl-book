package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
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
var out io.Writer = os.Stdout

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attrs := [][]byte{[]byte("")}
		for _, a := range n.Attr {
			b := []byte(fmt.Sprintf("%s=\"%s\"", a.Key, a.Val))
			attrs = append(attrs, b)
		}
		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s<%s%s>\n", depth*2, "", n.Data, bytes.Join(attrs, []byte(" ")))
		} else {
			fmt.Fprintf(out, "%*s<%s%s />\n", depth*2, "", n.Data, bytes.Join(attrs, []byte(" ")))
		}
		depth++
	}

	if n.Type == html.TextNode {
		b := []byte(n.Data)
		for _, l := range bytes.Split(b, []byte("\n")) {
			l = bytes.TrimSpace(l)
			if len(l) != 0 {
				fmt.Fprintf(out, "%*s%s\n", depth*2, "", l)
			}
		}
	}

	if n.Type == html.CommentNode {
		b := []byte(n.Data)
		lines := bytes.Split(b, []byte("\n"))
		if len(lines) == 1 {
			fmt.Fprintf(out, "%*s<!-- %s -->\n", depth*2, "", bytes.TrimSpace(lines[0]))
		} else {
			fmt.Fprintf(out, "%*s<!--\n", depth*2, "")
			for _, l := range lines {
				l = bytes.TrimSpace(l)
				if len(l) != 0 {
					fmt.Fprintf(out, "%*s%s\n", (depth+1)*2, "", l)
				}
			}
			fmt.Fprintf(out, "%*s-->\n", depth*2, "")
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
