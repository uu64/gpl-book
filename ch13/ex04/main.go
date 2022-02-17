package main

import (
	"bytes"
	"log"
	"os"

	"github.com/uu64/gpl-book/ch13/ex04/bzip"
)

func main() {
	var compressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	w.Write(b)

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	os.WriteFile("sampledata.dat.bz2", compressed.Bytes(), 0644)
}
