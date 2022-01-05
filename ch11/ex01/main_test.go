package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"testing"
)

const input = `100
Hello
世界
aaa
 

`

func TestCharCount(t *testing.T) {
	inr, inw, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	defer inr.Close()

	outr, outw, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	defer outr.Close()

	errr, errw, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	defer errr.Close()

	inw.Write([]byte(input))
	inw.Close()

	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw

	main()
	outw.Close()
	errw.Close()

	outbuf, err := ioutil.ReadAll(outr)
	if err != nil {
		log.Fatal(err)
	}

	errbuf, err := ioutil.ReadAll(errr)
	if err != nil {
		log.Fatal(err)
	}

	pat := []string{
		"'1'\t1",
		"'0'\t2",
		"'H'\t1",
		"'e'\t1",
		"'l'\t2",
		"'o'\t1",
		"'世'\t1",
		"'界'\t1",
		"'a'\t3",
		"' '\t1",
		"'\\\\n'\t6",
		"1\t18",
		"2\t0",
		"3\t2",
		"4\t0",
	}

	for _, p := range pat {
		matched, _ := regexp.Match(p, outbuf)
		if !matched {
			t.Errorf("'%s' should be matched.", p)
			t.Errorf("stdout: %s", outbuf)
			t.Errorf("stderr: %s", errbuf)
		}
	}
}
