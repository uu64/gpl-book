package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func createDir(path string) error {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		return err
	}
	return nil
}

func download(dirPath, fileName, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func parseURL(u *url.URL) (basename, dirPath string) {
	urlPath := fmt.Sprintf("%s%s", u.Hostname(), u.Path)
	elem := strings.Split(urlPath, "/")

	if elem[len(elem)-1] == "" {
		dirPath = urlPath
		basename = "index.html"
	} else {
		dirPath = strings.Join(elem[:len(elem)-1], "/")
		basename = elem[len(elem)-1]
		if path.Ext(basename) == "" {
			basename += ".html"
		}
	}
	return
}

func main() {
	downloaded := make(map[string]bool)

	crawl := func(rawURL string) []string {
		fmt.Println(rawURL)

		list := []string{}
		pageURL, err := url.Parse(rawURL)
		if err != nil {
			log.Print(err)
			return list
		}

		links, err := links.Extract(rawURL)
		if err != nil {
			log.Print(err)
			return list
		}
		for _, link := range links {
			l, err := url.Parse(link)
			if err != nil {
				log.Print(err)
				continue
			}

			if l.Hostname() == pageURL.Hostname() {
				basename, dirPath := parseURL(l)
				key := fmt.Sprintf("%s/%s", dirPath, basename)
				if downloaded[key] {
					continue
				}
				downloaded[key] = true

				if err := createDir(dirPath); err != nil {
					log.Print(err)
					continue
				}

				fmt.Printf("download from %s -> %s\n", l, path.Join(dirPath, basename))
				if err := download(dirPath, basename, link); err != nil {
					log.Print(err)
					continue
				}
				list = append(list, link)
			}
		}
		return list
	}
	breadthFirst(crawl, os.Args[1:])
}
