package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var opt int

func init() {
	flag.IntVar(&opt, "opt", 256, "Hash algorithm")
}

func main() {
	flag.Parse()
	if opt != 256 && opt != 384 && opt != 512 {
		fmt.Println("available option: 256 384 512")
		os.Exit(1)
	}

	var input string
	for {
		fmt.Scan(&input)
		switch opt {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
		}
	}
}
