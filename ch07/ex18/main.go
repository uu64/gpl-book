package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := parse(os.Stdin, os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		os.Exit(1)
	}
	printTree(n, 0)
}

func printTree(n Node, indent int) {
	switch n := n.(type) {
	case CharData:
		s := strings.TrimSpace(string(n))
		if len(s) > 0 {
			fmt.Printf("%s%s\n", strings.Repeat("  ", indent), s)
		}
	case *Element:
		items := []string{n.Type.Local}
		for _, a := range n.Attr {
			items = append(items, fmt.Sprintf("%s='%s'", a.Name.Local, a.Value))
		}
		fmt.Printf("%s<%s>\n", strings.Repeat("  ", indent), strings.Join(items, " "))
		for _, c := range n.Children {
			printTree(c, indent+1)
		}
		fmt.Printf("%s</%s>\n", strings.Repeat("  ", indent), n.Type.Local)
	}
}

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func parse(r io.Reader, w io.Writer, path []string) (Node, error) {
	dec := xml.NewDecoder(r)
	dummy := Element{xml.Name{}, []xml.Attr{}, []Node{}}
	stack := []*Element{&dummy}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			parent := stack[len(stack)-1]
			elem := Element{tok.Name, tok.Attr, []Node{}}
			parent.Children = append(parent.Children, &elem)
			stack = append(stack, &elem) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] //pop
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return dummy.Children[0], nil
}
