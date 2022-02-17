package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

type Decoder struct {
	lex *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	lex.next() // get the first token

	return &Decoder{lex}
}

type Token interface{}
type Delim rune

func (d Delim) String() string {
	return string(d)
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	lex := dec.lex
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("decode error at %s: %v", lex.scan.Position, x)
		}
	}()

	read(lex, reflect.ValueOf(v).Elem())
	return nil
}

func (dec *Decoder) More() bool {
	if dec.lex.token == scanner.EOF {
		return false
	}
	text := dec.lex.text()
	return text != ")"
}

func (dec *Decoder) Token() (Token, error) {
	defer dec.lex.next()
	return get(dec.lex)
}

// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func (dec *Decoder) Unmarshal(out interface{}) (err error) {
	// NOTE: Decodeと同じ実装になっているが問題の意図どおりか
	lex := dec.lex
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("unmarshal error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}
