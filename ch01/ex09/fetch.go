package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var httpScheme = "http://"
var httpsScheme = "https://"
var stdout = os.Stdout
var stderr = os.Stderr

func get(url string) error {
	if !strings.HasPrefix(url, httpScheme) && !strings.HasPrefix(url, httpsScheme) {
		url = httpScheme + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Fprintf(stdout, "response status code: %d\n", resp.StatusCode)
	fmt.Fprint(stdout, "response body:\n")
	_, err = io.Copy(stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("copying response body: %v", err)
	}

	return nil
}

func main() {
	for _, url := range os.Args[1:] {
		err := get(url)
		if err != nil {
			fmt.Fprintf(stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	}
}
