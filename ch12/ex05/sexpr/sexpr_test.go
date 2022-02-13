package sexpr

import (
	"encoding/json"
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
func Test(t *testing.T) {
	var str string
	sentence := "sample text"

	var numInt int
	size := 432

	var numFloat float32
	rate := 3.472

	var boolean bool
	flag := true

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
		t.Logf("Marshal() = %s\n", data)

		// can decode using json.Unmarshal
		err = json.Unmarshal(data, test.decoded)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		t.Logf("Unmarshal: %v\n", test.decoded)
	}
}
