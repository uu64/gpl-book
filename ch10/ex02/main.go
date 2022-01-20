package main

import (
	"fmt"
	"log"
	"os"

	"github.com/uu64/gpl-book/ch10/ex02/archive"
	_ "github.com/uu64/gpl-book/ch10/ex02/archive/tar"
	_ "github.com/uu64/gpl-book/ch10/ex02/archive/zip"
)

func main() {
	names, err := archive.Read(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}
