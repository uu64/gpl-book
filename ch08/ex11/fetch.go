package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func request(url string, done <-chan struct{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//lint:ignore SA1019 req.Cancel is required in this exercise.
	req.Cancel = done

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func fetch(urls []string) (url, filename string, n int64, err error) {
	if len(urls) == 0 {
		return
	}

	done := make(chan struct{})
	responses := make(chan *http.Response, len(url))
	for _, u := range urls {
		go func(url string) {
			res, err := request(url, done)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			responses <- res
		}(u)
	}
	res := <-responses

	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	close(done)

	local := path.Base(res.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return res.Request.URL.String(), "", 0, err
	}
	n, err = io.Copy(f, buf)

	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return res.Request.URL.String(), local, n, err
}

func main() {
	url, local, n, err := fetch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
		return
	}
	fmt.Fprintf(os.Stdout, "%s => %s (%d bytes).\n", url, local, n)
}
