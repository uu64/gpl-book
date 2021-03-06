package main

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func createURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func TestParseQuery(t *testing.T) {
	var tests = []struct {
		input *url.URL
		want  svgParams
	}{
		{createURL("http://localhost"), svgParams{600, 320, "white"}},
		{createURL("http://localhost?width=500"), svgParams{500, 320, "white"}},
		{createURL("http://localhost?height=300"), svgParams{600, 300, "white"}},
		{createURL("http://localhost?color=red"), svgParams{600, 320, "red"}},
		{createURL("http://localhost?color=%23ff0000ff"), svgParams{600, 320, "#ff0000ff"}},
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
