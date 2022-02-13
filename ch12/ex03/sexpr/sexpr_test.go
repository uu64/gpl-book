// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
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
func TestEncodeStruct1(t *testing.T) {
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

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEncodeStruct2(t *testing.T) {
	type Complex struct {
		Re  float32
		Im  float32
		Cpl complex64
	}
	c := Complex{
		Re:  3.0,
		Im:  8.2,
		Cpl: complex(3, 8),
	}

	// Encode it
	data, err := Marshal(c)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var cpl Complex
	if err := Unmarshal(data, &cpl); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", cpl)

	// Check equality.
	if !reflect.DeepEqual(cpl, c) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEncodeFloat(t *testing.T) {
	// Encode it
	data, err := Marshal(float32(2.3))
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var f float32
	if err := Unmarshal(data, &f); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", f)

	// Check equality.
	if float32(2.3) != f {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEncodeComplex(t *testing.T) {
	// Encode it
	data, err := Marshal(complex(2.3, 5))
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var cpl complex128
	if err := Unmarshal(data, &cpl); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", cpl)

	// Check equality.
	if complex(2.3, 5) != cpl {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(cpl)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEncodeBool1(t *testing.T) {
	// Encode it
	data, err := Marshal(true)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var b bool
	if err := Unmarshal(data, &b); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", b)

	// Check equality.
	if true != b {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEncodeBool2(t *testing.T) {
	// Encode it
	data, err := Marshal(nil)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var b bool
	if err := Unmarshal(data, &b); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", b)

	// Check equality.
	if false != b {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(b)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEncodeInterface(t *testing.T) {
	type Res struct {
		Data interface{}
	}
	res := Res{struct {
		Name, Message string
	}{"taro", "hello"}}

	// Encode it
	data, err := Marshal(res)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// NOTE: interface のデコーダは12.10で実装する
	// // Decode it
	// var r Res
	// if err := Unmarshal(data, &r); err != nil {
	// 	t.Fatalf("Unmarshal failed: %v", err)
	// }
	// t.Logf("Unmarshal() = %+v\n", r)

	// // Check equality.
	// if !reflect.DeepEqual(res, r) {
	// 	t.Fatal("not equal")
	// }

	// // Pretty-print it:
	// data, err = MarshalIndent(r)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf("MarshalIdent() = %s\n", data)
}
