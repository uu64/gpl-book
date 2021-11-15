package main

import (
	"fmt"

	"github.com/uu64/gpl-book/ch07/ex14/eval"
)

func main() {
	env := eval.Env{"x": 3, "y": 4}
	expr1, _ := eval.Parse("pow(x, 2) + pow(y, 2)")
	got1 := expr1.Eval(env)
	fmt.Println(got1)
	fmt.Println(expr1)

	expr2, _ := eval.Parse("+ 3.14159265358979")
	got2 := expr2.Eval(env)
	fmt.Println(got2)
	fmt.Println(expr2)

	expr3, _ := eval.Parse("5 / 9 * (F * 5 - 32)")
	got3 := expr2.Eval(env)
	fmt.Println(got3)
	fmt.Println(expr3)
}
