package main

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func createUrl(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func TestParseQuery(t *testing.T) {
	var tests = []struct {
		input *url.URL
		want  SvgParams
	}{
		{createUrl("http://localhost"), SvgParams{600, 320, "white"}},
		{createUrl("http://localhost?width=500"), SvgParams{500, 320, "white"}},
		{createUrl("http://localhost?height=300"), SvgParams{600, 300, "white"}},
		{createUrl("http://localhost?color=red"), SvgParams{600, 320, "red"}},
		{createUrl("http://localhost?color=%23ff0000ff"), SvgParams{600, 320, "#ff0000ff"}},
	}
	for _, test := range tests {
		got, err := parseQuery(test.input)
		fmt.Println(got)
		if err != nil {
			t.Errorf("parseQuery(%v) failed: %v\n", test.input, err)
		}
		if got != test.want {
			t.Errorf("parseQuery(%v) = %v\n", test.input, got)
		}
	}

}
