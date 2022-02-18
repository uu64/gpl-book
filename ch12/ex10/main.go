package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/uu64/gpl-book/ch12/ex10/sexpr"
)

func main() {
	type Sample struct {
		NumInt   int
		NumFloat float32
		Str      string
		Flag     bool
		Cmplx    complex64
	}

	data := Sample{34, 2.3, "hello", true, complex(1, 2.3)}
	b, _ := sexpr.Marshal(data) // ignoring errors
	fmt.Println(string(b))

	var sample Sample
	stream := `
	((NumInt 34)
	 (NumFloat 2.300000)
	 (Str "hello")
	 (Flag t)
	 (Cmplx #C(1.000000 2.300000)))
	`

	dec := sexpr.NewDecoder(strings.NewReader(stream))
	err := dec.Unmarshal(&sample)

	if err != nil {
		log.Fatalf("Decode failed: %v", err)
	}
	fmt.Printf("Decode() = \n%v\n", sample)

	type Res struct {
		Data interface{}
	}
	r := Res{[]int{1, 2, 3}}
	b2, _ := sexpr.Marshal(r)
	fmt.Println(string(b2))

	dec2 := sexpr.NewDecoder(bytes.NewReader(b2))
	var res Res
	err = dec2.Unmarshal(&res)
	if err != nil {
		log.Fatalf("Decode failed: %v", err)
	}
	fmt.Printf("Decode() = \n%v\n", res)
}
