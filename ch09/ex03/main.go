package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/uu64/gpl-book/ch09/ex03/memo"
)

func httpGetBody(url string, cancelled ...<-chan struct{}) (interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// 先頭しか使わない
	if len(cancelled) > 0 {
		//lint:ignore SA1019 req.Cacnel is required in this exercise.
		req.Cancel = cancelled[0]
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := memo.New(httpGetBody)
	defer m.Close()

	url := "http://gopl.io"

	var n int
	for n < 2 {
		start := time.Now()
		done := make(chan struct{})
		res := make(chan interface{})

		go func() {
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
			}
			res <- value
		}()
		if n == 0 {
			// time.Sleep(300 * time.Millisecond)
			close(done)
		}

		// value := <-res
		// fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		<-res
		fmt.Printf("%s, %s\n", url, time.Since(start))
		n++
	}
}
