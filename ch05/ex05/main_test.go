package main

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
    <link rel="stylesheet" href="styles.css">
    <script src="javascript.js"></script>
  </head>
  <body>
    <h1>My First Heading</h1>
    <p>My first paragraph.</p>
    <h2>Absolute URLs</h2>
    <p><a href="https://www.w3.org/">w3c</a></p>
    <p><a href="https://www.google.com/">google</a></p>
    <h2>relative urls</h2>
    <a href="default.asp">
      <img src="smiley.gif" alt="HTML tutorial" style="width:42px;height:42px;">
    </a>
    <p><a href="/css/default.asp">css tutorial</a></p>
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

func TestCountWordsAndImages(t *testing.T) {
	n, err := html.Parse(strings.NewReader(doc))
	if err != nil {
		log.Fatal("fail: parse document")
	}
	// My, First, HTML, My, First, Heading, My, first, paragraph., Absolute, URLs,
	// w3c, google, relative, urls, css, tutorial, An, Unordered, HTML, List,
	// Coffee, Tea, Milk, An, Ordered, HTML, List, Coffee, Tea, Milk,
	words := 31
	// smiley.gif,
	images := 1
	w, i := countWordsAndImages(n)
	if w != words || i != images {
		t.Errorf("fail: got -> words: %d, images: %d\n", w, i)
	}

}
