package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	parse(os.Stdin, os.Stdout, os.Args[1:])
}

func parse(r io.Reader, w io.Writer, path []string) {
	dec := xml.NewDecoder(r)
	var stack []*xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, &tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, path) {
				parents := []string{}
				for _, s := range stack {
					attrs := []string{}
					for _, a := range s.Attr {
						attrs = append(attrs, fmt.Sprintf("%s=%s", a.Name.Local, a.Value))
					}
					parents = append(parents, fmt.Sprintf("%s(%s)", s.Name.Local, strings.Join(attrs, ",")))
				}
				fmt.Fprintf(w, "%s: %s\n", strings.Join(parents, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []*xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].Name.Local == y[0] {
			y = y[1:]
		} else {
			for _, a := range x[0].Attr {
				if a.Value == y[0] {
					y = y[1:]
					break
				}
			}
		}
		x = x[1:]
	}
	return false
}
