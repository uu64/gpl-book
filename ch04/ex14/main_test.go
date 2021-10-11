package main

import (
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

func TestParseURL(t *testing.T) {
	var tests = []struct {
		input   *url.URL
		owner   string
		repo    string
		message string
	}{
		{createURL("http://localhost/issues"), "", "", "error: please set the owner and repo in url path"},
		{createURL("http://localhost/issues/uu64/gpl-book"), "uu64", "gpl-book", ""},
		{createURL("http://localhost/issues/uu64/gpl-book"), "uu64", "gpl-book", ""},
		{createURL("http://localhost/issues/uu64/gpl-book/"), "uu64", "gpl-book", ""},
	}
	for _, test := range tests {
		owner, repo, err := parseURL(test.input)
		if err != nil {
			if err.Error() != test.message {
				t.Errorf("parseQuery(%v) = %s, %s, %s\n", test.input, owner, repo, err.Error())
			}
		} else {
			if owner != test.owner || repo != test.repo {
				t.Errorf("parseQuery(%v) = %s, %s, %s\n", test.input, owner, repo, err.Error())
			}
		}
	}

}

func TestParseQuery(t *testing.T) {
	var tests = []struct {
		input *url.URL
		want  string
	}{
		{createURL("http://localhost/issues/uu64/gpl-book"), ""},
		{createURL("http://localhost/issues/uu64/gpl-book?is=open"), "open"},
		{createURL("http://localhost/issues/uu64/gpl-book/?is=closed"), "closed"},
	}
	for _, test := range tests {
		got := parseQuery(test.input)
		if got != test.want {
			t.Errorf("parseQuery(%v) = %v\n", test.input, got)
		}
	}

}
