package showtext

import (
	"bytes"
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

const want = `My First HTML
My First Heading
My first paragraph.
An Unordered HTML List
Coffee
Tea
Milk
An Ordered HTML List
Coffee
Tea
Milk
`

func TestShowText(t *testing.T) {
	n, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		log.Fatal("fail: parse document")
	}
	buf := &bytes.Buffer{}
	ShowText(buf, n)
	if buf.String() != want {
		t.Errorf("fail:\n%v\n", buf.String())
	}
}
