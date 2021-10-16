package countelem

import (
	"log"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

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

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func TestCountElem(t *testing.T) {
	n, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		log.Fatal("fail: parse document")
	}
	counter := make(map[string]int)
	want := map[string]int{
		"html":  1,
		"head":  1,
		"title": 1,
		"meta":  1,
		"body":  1,
		"h1":    1,
		"h2":    2,
		"p":     1,
		"ul":    1,
		"ol":    1,
		"li":    6,
	}
	CountElem(counter, n)
	if !equal(counter, want) {
		t.Errorf("fail: got %v\n", counter)
	}

}
