package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	p := printer{width: 0}
	if err := pretty(&p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

type Kind int

const (
	kVal Kind = iota
	kBlank
	kStart
	kEnd
	kBr
)

type token struct {
	kind Kind // one of "s ()" (string, blank, start, end)
	str  string
}

type printer struct {
	tokens []*token // FIFO buffer
	stack  []*token // stack of open ' ' and '(' tokens
	rtotal int      // total number of spaces needed to print stream

	bytes.Buffer
	indents []int
	width   int // width of the string
}

func (p *printer) string(str string) {
	tok := &token{kind: kVal, str: str}
	if len(p.stack) == 0 {
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotal += len(str)
	}
}
func (p *printer) pop() (top *token) {
	last := len(p.stack) - 1
	top, p.stack = p.stack[last], p.stack[:last]
	return
}
func (p *printer) begin() {
	if len(p.stack) == 0 {
		p.rtotal = 1
	}
	t := &token{kind: kStart}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t) // push
	p.string("(")
}
func (p *printer) end() {
	p.string(")")
	p.tokens = append(p.tokens, &token{kind: kEnd})
	x := p.pop()
	if x.kind == kBlank || x.kind == kBr {
		p.pop()
	}
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}
func (p *printer) space() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == kBlank {
		p.stack = p.stack[:last] // pop
	}
	t := &token{kind: kBlank}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}
func (p *printer) lineBreak() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == kBr {
		p.stack = p.stack[:last] // pop
	}
	t := &token{kind: kBr}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}
func (p *printer) print(t *token) {
	// fmt.Printf("%d, %d, %s, %d\n", p.indents, t.kind, t.str, p.width)
	switch t.kind {
	case kVal:
		p.WriteString(t.str)
		p.width += len(t.str)
	case kBr:
		p.width = p.indents[len(p.indents)-1] + 1
		fmt.Fprintf(&p.Buffer, "\n%*s", p.width, "")
	case kStart:
		p.indents = append(p.indents, p.width)
	case kEnd:
		p.indents = p.indents[:len(p.indents)-1] // pop
	case kBlank:
		p.WriteByte(' ')
		p.width++
	}
}
func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p.stringf("%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		p.stringf("%f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		p.stringf("#C(%f %f)", real(c), imag(c))

	case reflect.Bool:
		if v.Bool() {
			p.string("t")
		} else {
			p.string("nil")
		}

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Array, reflect.Slice: // (value ...)
		p.begin()
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				p.lineBreak()
			}
			if err := pretty(p, v.Index(i)); err != nil {
				return err
			}
		}
		p.end()

	case reflect.Struct: // ((name value ...)
		p.begin()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				p.lineBreak()
			}
			p.begin()
			p.string(v.Type().Field(i).Name)
			p.space()
			if err := pretty(p, v.Field(i)); err != nil {
				return err
			}
			p.end()
		}
		p.end()

	case reflect.Map: // ((key value ...)
		p.begin()
		for i, key := range v.MapKeys() {
			if i > 0 {
				p.lineBreak()
			}
			p.begin()
			if err := pretty(p, key); err != nil {
				return err
			}
			p.space()
			if err := pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
			p.end()
		}
		p.end()

	case reflect.Ptr:
		return pretty(p, v.Elem())

	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
