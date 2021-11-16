package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/uu64/gpl-book/ch07/ex15/eval"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Fprintf(os.Stderr, "internal error: %v\n", p)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	// read a single expression
	fmt.Printf("expr?: ")
	scanner.Scan()
	input := scanner.Text()
	// parse
	expr, err := eval.Parse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse the input: %v\n", err)
		os.Exit(1)
	}

	// get variables and handle static errors
	vars := make(map[eval.Var]bool)
	if err = expr.Check(vars); err != nil {
		fmt.Fprintf(os.Stderr, "invalid expr: %v\n", err)
		os.Exit(1)
	}

	// prompts the user to provide values
	env := eval.Env{}
	for k := range vars {
		fmt.Printf("'%s'?: ", k)
		for scanner.Scan() {
			s := scanner.Text()
			v, err := strconv.ParseFloat(s, 64)
			if err == nil {
				env[k] = v
				break
			} else {
				fmt.Println("Variable must be a number. Please type it again.")
				fmt.Printf("'%s'?: ", k)
			}
		}
	}

	fmt.Printf("\nanswer: %f\n", expr.Eval(env))

	fmt.Println("finish.")
}
