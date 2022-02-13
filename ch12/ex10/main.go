package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/uu64/gpl-book/ch12/ex10/sexpr"
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

	var movie Movie
	// sexpr.Unmarshal(b, &movie) // ignore an error
	// fmt.Println(movie)
	// fmt.Println()

	decs := sexpr.NewDecoder(bytes.NewReader(b))
	for decs.More() {
		// t, err := decs.Token()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%T: %v\n", t, t)

		err := decs.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(movie)
	}
	fmt.Println()

	type Message struct {
		Name, Text string
	}

	const sStream = `
	(((Name "Ed")
	  (Text "Knock knock."))
	 ((Name "Sam")
	  (Text "Who's there?"))
	 ((Name "Ed")
	  (Text "Go fmt."))
	 ((Name "Sam")
	  (Text "Go fmt who?"))
	 ((Name "Ed")
	  (Text "Go fmt yourself!")))
	`
	decs = sexpr.NewDecoder(strings.NewReader(sStream))
	// read open bracket
	t, err := decs.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	for decs.More() {
		// t, err := decs.Token()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%T: %v\n", t, t)

		var m Message
		err := decs.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}
	// read closing bracket
	t, err = decs.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	fmt.Println()

	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
	`
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		// t, err := dec.Token()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%T: %v\n", t, t)

		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(m)
		// fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}
