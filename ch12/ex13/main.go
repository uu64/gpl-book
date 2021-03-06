package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/uu64/gpl-book/ch12/ex13/sexpr"
)

func main() {
	type Sample struct {
		NumInt   int     `sexpr:"int"`
		NumFloat float32 `sexpr:"float"`
		Str      string  `sexpr:"s"`
		Flag     bool    `sexpr:"f"`
		Cmplx    complex64
	}

	data := Sample{34, 2.3, "hello", true, complex(1, 2.3)}
	b, _ := sexpr.Marshal(data) // ignoring errors
	fmt.Println(string(b))

	var sample Sample
	// stream := `
	// ((NumInt 34)
	//  (NumFloat 2.300000)
	//  (Str "hello")
	//  (Flag t)
	//  (Cmplx #C(1.000000 2.300000)))
	// `
	stream := `
	((int 34)
	 (float 2.300000)
	 (s "hello")
	 (f t)
	 (Cmplx #C(1.000000 2.300000)))
	`

	dec := sexpr.NewDecoder(strings.NewReader(stream))
	err := dec.Unmarshal(&sample)

	if err != nil {
		log.Fatalf("Decode failed: %v", err)
	}
	fmt.Printf("Decode() = \n%v\n", sample)
}
