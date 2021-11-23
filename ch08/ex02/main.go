package main

import (
	"log"

	"github.com/uu64/gpl-book/ch08/ex02/ftp"
)

func main() {
	if err := ftp.ListenAndServe(":21"); err != nil {
		log.Fatal(err)
	}
}
