package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"text/template"

	"golang.org/x/net/html"
)

var tmpl = template.Must(template.New("tmpl").Parse(`<html>
  <head>
    <title>
      Sample document
    </title>
  </head>
  <body>
    {{.}}
  </body>
</html>
`))

func TestForEachNode(t *testing.T) {
	tests := []struct {
		doc  string
		want string
	}{
		{
			buildHTML(`<img src="apple.gif" alt="This is an apple" style="width:42px;height:42px;"></img>`),
			buildHTML(`<img src="apple.gif" alt="This is an apple" style="width:42px;height:42px;" />`),
		},
		{
			buildHTML("<p>This is a sample paragraph.</p>"),
			buildHTML("<p>\n      This is a sample paragraph.\n    </p>"),
		},
		{
			buildHTML("<!-- This is a sample comment. -->"),
			buildHTML("<!-- This is a sample comment. -->"),
		},
		{
			buildHTML("<!-- This is a \nsample comment. -->"),
			buildHTML("<!--\n      This is a\n      sample comment.\n    -->"),
		},
	}

	for _, test := range tests {
		out = new(bytes.Buffer)
		n, err := html.Parse(strings.NewReader(test.doc))
		if err != nil {
			log.Fatal("fail: parse document")
		}

		forEachNode(n, startElement, endElement)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("fail: got %v\n", got)
			t.Errorf("fail: want %v\n", test.want)
		}
	}
}

func buildHTML(body string) string {
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, body); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
