package main

import (
	"fmt"
	"log"

	"github.com/uu64/gpl-book/ch12/ex03/sexpr"
)

func main() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    true,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	b, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	type Complex struct {
		re float32
		im float32
		c  complex64
	}
	c := Complex{
		re: 3.0,
		im: 8.2,
		c:  complex(3.0, 8.2),
	}
	b2, err := sexpr.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b2))

	b3, err := sexpr.Marshal(complex(2, 3))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b3))

	type Res struct {
		Data interface{}
	}
	res := Res{c}
	b4, err := sexpr.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b4))

	res2 := Res{45}
	b5, err := sexpr.Marshal(res2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b5))

	res3 := Res{[]int{1, 2, 3}}
	b6, err := sexpr.Marshal(res3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b6))
}
