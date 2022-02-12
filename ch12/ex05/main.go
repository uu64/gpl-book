package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/uu64/gpl-book/ch12/ex05/sexpr"
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
	b, err := json.MarshalIndent(strangelove, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	b2, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b2))

	b3, err := json.MarshalIndent(34, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b3))

	b4, err := sexpr.Marshal([]string{
		"test",
		"hello",
		"world",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b4))

	b5, err := json.MarshalIndent([]string{
		"test",
		"hello",
		"world",
	}, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b5))

	count2 := map[int][]string{
		2: {"hello", "world"},
		3: {"test", "text", "word"},
	}
	b6, err := sexpr.Marshal(count2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b6))

	b7, err := json.Marshal(count2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b7))
}
