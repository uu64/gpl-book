package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
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
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
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

func get(lex *lexer) (Token, error) {
	switch lex.token {
	case scanner.Ident:
		// fmt.Printf("ident: %s\n", lex.text())
		if lex.text() == "t" {
			return true, nil
		} else if lex.text() == "nil" {
			return false, nil
		}
		return lex.text(), nil
	case scanner.String:
		// fmt.Printf("string: %s\n", lex.text())
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		return s, nil
	case scanner.Int:
		// fmt.Println("int")
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		return i, nil
	case scanner.Float:
		// fmt.Println("float")
		i, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring errors
		return i, nil
	case '(', ')', '#':
		// fmt.Println(lex.token)
		return Delim(lex.token), nil
	case scanner.EOF:
		// fmt.Println("EOF")
		return nil, fmt.Errorf("EOF")
	}
	return nil, fmt.Errorf("unexpected token %q", lex.text())
}
