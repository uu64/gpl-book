package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
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
	kSep
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
	indent int
	width  int // width of the string
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
}
func (p *printer) end() {
	p.tokens = append(p.tokens, &token{kind: kEnd})
	x := p.pop()
	if x.kind == kSep || x.kind == kBr {
		p.pop()
	}
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}
func (p *printer) sep() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == kSep {
		p.stack = p.stack[:last] // pop
	}
	t := &token{kind: kSep}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}
func (p *printer) lineBreak() {
	t := &token{kind: kBr}
	p.tokens = append(p.tokens, t)
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}
func (p *printer) print(t *token) {
	// fmt.Printf("%d, %d, %s, %d\n", p.indents, t.kind, t.str, p.width)
	switch t.kind {
	case kVal:
		// fmt.Printf("val: %s\n", t.str)
		p.WriteString(t.str)
		p.width += len(t.str)
	case kBr:
		// fmt.Println("br")
		fmt.Fprintf(&p.Buffer, "\n%*s", p.indent*2, "")
	case kStart:
		// fmt.Println("start")
		p.indent++
	case kEnd:
		// fmt.Println("end")
		p.indent--
	case kSep:
		// fmt.Println("sep")
		p.WriteString(": ")
		p.width += 2
	}
}
func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p.stringf("%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		p.stringf("%f", v.Float())

	case reflect.Bool:
		p.stringf("%t", v.Bool())

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Array, reflect.Slice: // (value ...)
		p.begin()
		p.string("[")
		p.lineBreak()
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				p.string(",")
				p.lineBreak()
			}
			if err := pretty(p, v.Index(i)); err != nil {
				return err
			}
		}
		p.end()
		p.lineBreak()
		p.string("]")

	case reflect.Struct: // ((name value ...)
		p.begin()
		p.string("{")
		p.lineBreak()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				p.string(",")
				p.lineBreak()
			}
			p.stringf("%q", v.Type().Field(i).Name)
			p.sep()
			if err := pretty(p, v.Field(i)); err != nil {
				return err
			}
		}
		p.end()
		p.lineBreak()
		p.string("}")

	case reflect.Map: // ((key value ...)
		p.begin()
		p.string("{")
		p.lineBreak()
		for i, key := range v.MapKeys() {
			if i > 0 {
				p.string(",")
				p.lineBreak()
			}
			switch key.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				p.stringf("%q", strconv.FormatInt(key.Int(), 10))

			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				p.stringf("%q", strconv.FormatUint(key.Uint(), 10))

			case reflect.String:
				p.stringf("%q", key.String())

			default: // the key must either be a string or an integer type
				return fmt.Errorf("unsupported type: %s", v.Type())
			}
			p.sep()
			if err := pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
		}
		p.end()
		p.lineBreak()
		p.string("}")

	case reflect.Ptr:
		return pretty(p, v.Elem())

	case reflect.Interface:
		return pretty(p, reflect.ValueOf(v.Interface()))

	default: // complex, map, chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
