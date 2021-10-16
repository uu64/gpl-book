package countelem

import "golang.org/x/net/html"

// CountElem counts the number of each element
func CountElem(counter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data] += 1
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CountElem(counter, c)
	}
}
