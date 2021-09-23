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
		want  imgParam
	}{
		{createURL("http://localhost"), imgParam{-2.0, -2.0, 1}},
		{createURL("http://localhost?x=0.5"), imgParam{0.5, -2.0, 1}},
		{createURL("http://localhost?y=0.5"), imgParam{-2.0, 0.5, 1}},
		{createURL("http://localhost?zoom=100"), imgParam{-2.0, -2.0, 100}},
		{createURL("http://localhost?x=-1&y=-0.5&zoom=100"), imgParam{-1.0, -0.5, 100}},
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
