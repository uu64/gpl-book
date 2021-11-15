package eval

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
	}{
		{"sqrt(A / pi)"},
		{"sqrt(A / 3.14)"},
		{"pow(x, 3) + pow(y, 3)"},
		{"5 / 9 * (F - 32)"},
		{"5 / 9 * (F + 3 - 32)"},
		{"5 / 9 * (F * 3 - 32)"},
		{"5 / 9 * (F - 3 / 32)"},
		{"5 / 9 * (F - 3 / 32 - 4)"},
		{"-1 + -x"},
		{"-1 - x"},
	}
	for _, test := range tests {
		expr1, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		expr2, err := Parse(expr1.String())
		if err != nil {
			t.Errorf("failed to reparse: %v", err) // parse error
		}

		if isEqual := reflect.DeepEqual(expr1, expr2); !isEqual {
			t.Errorf("2 trees did not match: %s, %s\n", expr1.String(), expr2.String())
			fmt.Println(expr1.Eval(Env{"F": 32}))
			fmt.Println(expr2.Eval(Env{"F": 32}))
		}
	}
}
