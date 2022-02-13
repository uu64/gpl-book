package sexpr

import (
	"strings"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func TestMarshal(t *testing.T) {
	var str string
	sentence := "sample text"

	var numInt int
	size := 432

	var numFloat float32
	rate := 3.472

	var boolean bool
	flag := true

	var cmplx complex128
	cmp := complex128(complex(1, 2.5))

	var array []string
	words := []string{
		"hello",
		"world",
		"!!!",
	}

	var dict map[string]int
	count := map[string]int{
		"hello": 1,
		"world": 15,
		"!!!":   2,
	}

	var dictarray map[int][]string
	count2 := map[int][]string{
		1:  {"hello", "world"},
		4:  {"test", "text", "word"},
		23: {"a"},
	}

	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	var movie Movie
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

	type Res struct {
		Data interface{}
	}
	var res Res
	movieRes := Res{strangelove}

	tests := []struct {
		encoded interface{}
		decoded interface{}
	}{
		{sentence, &str},
		{size, &numInt},
		{rate, &numFloat},
		{flag, &boolean},
		{cmp, &cmplx},
		{words, &array},
		{count, &dict},
		{count2, &dictarray},
		{strangelove, &movie},
		{movieRes, &res},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.encoded)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		t.Logf("Marshal() = \n%s\n", data)
	}
}

func TestUnmarshal(t *testing.T) {
	var str string
	var numInt int
	var numFloat float32
	var boolean bool
	var cmplx complex128
	var array []string
	var dict map[string]int
	var dictarray map[int][]string
	var movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	var res struct {
		Data interface{}
	}

	tests := []struct {
		encoded string
		decoded interface{}
	}{
		{"\"hello\"", &str},
		{"42", &numInt},
		{"0.576", &numFloat},
		{"t", &boolean},
		{"nil", &boolean},
		{"#C(1.0 2.0)", &cmplx},
		{"(\"hello\" \"world\" \"!!!\")", &array},
		{`(
			("hello" 1)
			("world" 15)
			("!!!" 2)
		)`, &dict},
		{`(
			(1 ("hello" "world"))
			(4 ("test" "text" "word"))
			(23 ("a"))
		)`, &dictarray},
		{`
		((Title "Dr. Strangelove")
		(Subtitle "How I Learned to Stop Worrying and Love the Bomb")
		(Year 1964)
		(Color t)
		(Actor (("Dr. Strangelove" "Peter Sellers")
				("Grp. Capt. Lionel Mandrake" "Peter Sellers")
				("Pres. Merkin Muffley" "Peter Sellers")
				("Gen. Buck Turgidson" "George C. Scott")
				("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
				("Maj. T.J. \"King\" Kong" "Slim Pickens")))
		(Oscars ("Best Actor (Nomin.)"
				"Best Adapted Screenplay (Nomin.)"
				"Best Director (Nomin.)"
				"Best Picture (Nomin.)")))
		`, &movie},
		{`
		((Data 34))
		`, &res},
		{`
		((Data ((Title "Dr. Strangelove")
				(Subtitle "How I Learned to Stop Worrying and Love the Bomb")
				(Year 1964)
				(Color t)
				(Actor (("Dr. Strangelove" "Peter Sellers")
						("Grp. Capt. Lionel Mandrake" "Peter Sellers")
						("Pres. Merkin Muffley" "Peter Sellers")
						("Gen. Buck Turgidson" "George C. Scott")
						("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
						("Maj. T.J. \"King\" Kong" "Slim Pickens")))
				(Oscars ("Best Actor (Nomin.)"
						"Best Adapted Screenplay (Nomin.)"
						"Best Director (Nomin.)"
						"Best Picture (Nomin.)")))))
		`, &res},
	}

	for _, test := range tests {
		dec := NewDecoder(strings.NewReader(test.encoded))
		// Decode it
		err := dec.Unmarshal(test.decoded)
		if err != nil {
			t.Fatalf("Decode failed: %v", err)
		}
		t.Logf("Decode() = \n%v\n", test.decoded)
	}
}

func TestStream(t *testing.T) {
	var str string
	var numInt int
	var numFloat float32
	var boolean bool
	var cmplx complex128
	var array []string
	var dict map[string]int
	var dictarray map[int][]string
	var movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	var message struct {
		Name, Text string
	}

	tests := []struct {
		encoded string
		decoded interface{}
	}{
		{"\"hello\"", &str},
		{"42", &numInt},
		{"0.576", &numFloat},
		{"t", &boolean},
		{"nil", &boolean},
		{"#C(1.0 2.0)", &cmplx},
		{"(\"hello\" \"world\" \"!!!\")", &array},
		{`(
			("hello" 1)
			("world" 15)
			("!!!" 2)
		)`, &dict},
		{`(
			(1 ("hello" "world"))
			(4 ("test" "text" "word"))
			(23 ("a"))
		)`, &dictarray},
		{`
		((Title "Dr. Strangelove")
		(Subtitle "How I Learned to Stop Worrying and Love the Bomb")
		(Year 1964)
		(Color t)
		(Actor (("Dr. Strangelove" "Peter Sellers")
				("Grp. Capt. Lionel Mandrake" "Peter Sellers")
				("Pres. Merkin Muffley" "Peter Sellers")
				("Gen. Buck Turgidson" "George C. Scott")
				("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
				("Maj. T.J. \"King\" Kong" "Slim Pickens")))
		(Oscars ("Best Actor (Nomin.)"
				"Best Adapted Screenplay (Nomin.)"
				"Best Director (Nomin.)"
				"Best Picture (Nomin.)")))
		`, &movie},
		{`
		((Name "Ed")
		 (Text "Knock knock."))
		((Name "Sam")
		 (Text "Who's there?"))
		((Name "Ed")
		 (Text "Go fmt."))
		((Name "Sam")
		 (Text "Go fmt who?"))
		((Name "Ed")
		 (Text "Go fmt yourself!"))
		`, &message},
	}

	for _, test := range tests {
		dec := NewDecoder(strings.NewReader(test.encoded))
		for dec.More() {
			// Decode it
			err := dec.Decode(test.decoded)
			if err != nil {
				t.Fatalf("Decode failed: %v", err)
			}
			t.Logf("Decode() = \n%v\n", test.decoded)
		}
	}
}
