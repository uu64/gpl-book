// Package sexpr provides a means for converting Go objects to and
// from S-expressions.
package sexpr

import (
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if lex.text() == "nil" && v.IsValid() {
			v.Set(reflect.Zero(v.Type()))
		}
		if lex.text() == "t" && v.IsValid() {
			v.Set(reflect.ValueOf(true))
		}
		lex.next()
		return
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		if v.IsValid() {
			v.Set(reflect.ValueOf(s))
		}
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		if v.IsValid() {
			switch v.Kind() {
			case reflect.Interface:
				v.Set(reflect.ValueOf(i))
			default:
				v.SetInt(int64(i))
			}
		}
		lex.next()
		return
	case scanner.Float:
		i, _ := strconv.ParseFloat(lex.text(), 64) // NOTE: ignoring errors
		if v.IsValid() {
			switch v.Kind() {
			case reflect.Interface:
				v.Set(reflect.ValueOf(i))
			default:
				v.SetFloat(i)
			}
		}
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	case '#':
		lex.next()
		lex.next() // consume 'C'
		lex.next() // consume '('
		re, _ := strconv.ParseFloat(lex.text(), 64)
		lex.next() // consume ' '
		im, _ := strconv.ParseFloat(lex.text(), 64)
		if v.IsValid() {
			switch v.Kind() {
			case reflect.Interface:
				v.Set(reflect.ValueOf(complex(re, im)))
			default:
				v.SetComplex(complex128(complex(re, im)))
			}
		}
		// readList(lex, v)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		fields := make(map[string]reflect.Value)
		for i := 0; i < v.NumField(); i++ {
			fieldInfo := v.Type().Field(i)
			tag := fieldInfo.Tag
			name := tag.Get("sexpr")
			if name == "" {
				name = fieldInfo.Name
			}
			fields[name] = v.Field(i)
		}

		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			f := fields[name]
			read(lex, f)
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	// TODO: 未実装
	// case reflect.Interface:

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Kind()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
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
