package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
	"text/template"

	"golang.org/x/net/html"
)

var tmpl = template.Must(template.New("tmpl").Parse(`<html>
  <head />
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
			`<img src="apple.gif" alt="This is an apple" style="width:42px;height:42px;"></img>`,
			buildHTML(`<img src="apple.gif" alt="This is an apple" style="width:42px;height:42px;" />`)},
		{
			"<p>This is a sample paragraph.</p>",
			buildHTML("<p>\n      This is a sample paragraph.\n    </p>")},
		{
			buildHTML("<!-- This is a sample commnt. -->"),
			buildHTML("<p>\n      This is a sample paragraph.\n    </p>")},
	}

	for _, test := range tests {
		out = new(bytes.Buffer)
		n, err := html.Parse(strings.NewReader(test.doc))
		if err != nil {
			log.Fatal("fail: parse document")
		}

		forEachNode(n, startElement, endElement)
		got := out.(*bytes.Buffer).String()
		fmt.Println(got)
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
