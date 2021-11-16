package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x.String())
}

func (b binary) String() string {
	var x, y string
	x = b.x.String()
	y = b.y.String()

	if string(b.op) == "*" || string(b.op) == "/" {
		switch b.x.(type) {
		case binary:
			if string(b.x.(binary).op) == "+" || string(b.x.(binary).op) == "-" {
				x = fmt.Sprintf("(%s)", x)
			}
		}
		switch b.y.(type) {
		case binary:
			if string(b.y.(binary).op) == "+" || string(b.y.(binary).op) == "-" {
				y = fmt.Sprintf("(%s)", y)
			}
		}
	}
	return fmt.Sprintf("%s %s %s", x, string(b.op), y)
}

func (c call) String() string {
	args := []string{}
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}

func (f factorial) String() string {
	return fmt.Sprintf("%s%s", f.x.String(), string(f.op))
}
