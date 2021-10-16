package showtext

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// ShowText shows the content of text nodes in html document
func ShowText(w io.Writer, n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		content := strings.TrimSpace(n.Data)
		if len(content) != 0 {
			fmt.Fprintf(w, "%s\n", content)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ShowText(w, c)
	}
}
