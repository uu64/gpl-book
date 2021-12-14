package main

import (
	"bytes"
	"flag"
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

type link struct {
	url   string
	depth int
}

var tokens = make(chan struct{}, 20)
var maxDepth = flag.Int("depth", 1, "Max depth for clawling.")
var mirrorURL = flag.String("mirror", "http://localhost:8080", "URL of the mirrored page.")

func toLinks(list []string, depth int) []link {
	links := []link{}
	for _, l := range list {
		links = append(links, link{l, depth})
	}
	return links
}

func crawl(url string, depth int) []link {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return toLinks(list, depth+1)
}
func createDir(path string) error {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		return err
	}
	return nil
}

func download(dirPath, fileName string, u *url.URL) error {
	// NOTE: 画像やJSファイルなどのアセットの取得は未対応
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return err
	}
	// ミラーページのURLに書き換え
	_, err = io.Copy(out, bytes.NewBuffer(bytes.ReplaceAll(buf.Bytes(), []byte(fmt.Sprintf("%s://%s", u.Scheme, u.Hostname())), []byte(*mirrorURL))))
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
	flag.Parse()

	worklist := make(chan []link)
	var n int // number of pending sends to worklist
	var depth int

	// Start with the command-line arguments.
	args := flag.Args()
	remote, err := url.Parse(args[0])
	if err != nil {
		log.Fatal(err)
	}

	n++
	go func() { worklist <- toLinks([]string{args[0]}, depth) }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		links := <-worklist
		for _, l := range links {
			u, err := url.Parse(l.url)
			if err != nil {
				log.Fatal(err)
			}

			// 別ドメインのページはスキップ
			if u.Hostname() != remote.Hostname() {
				continue
			}

			if !seen[l.url] && l.depth < *maxDepth {
				seen[l.url] = true
				n++

				go func(l link, u *url.URL) {
					// コンテンツをダウンロード
					baseName, dirPath := parseURL(u)
					if err := createDir(dirPath); err != nil {
						// ディレクトリ作成に失敗したらログ出力して諦める
						log.Print(err)
						return
					}
					fmt.Printf("download from %s -> %s\n", l.url, path.Join(dirPath, baseName))
					if err := download(dirPath, baseName, u); err != nil {
						log.Print(err)
						// ダウンロードに失敗したらログ出力して諦める
						return
					}

					// リンクを収集
					worklist <- crawl(l.url, l.depth)
				}(l, u)
			}
		}
	}
}
