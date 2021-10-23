package main

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var doc = `<html>
  <head>
    <title>
      Sample document
    </title>
  </head>
  <body>
    <p id="main">This is a sample document</p>
  </body>
</html>
`

func TestElementById(t *testing.T) {
	tests := []struct {
		id      string
		isFound bool
	}{
		{"main", true},
		{"aaa", false},
	}

	for _, test := range tests {
		n, err := html.Parse(strings.NewReader(doc))
		if err != nil {
			log.Fatal("fail: parse document")
		}

		ElementByID(n, test.id)
		if (targetNode != nil) != test.isFound {
			t.Errorf("fail: element with id '%s' -> %v\n", test.id, targetNode)
		}
	}
}
