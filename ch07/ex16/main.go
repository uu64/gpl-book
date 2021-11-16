package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/uu64/gpl-book/ch07/ex16/eval"
)

type reqBody struct {
	Expr string `json:"expr"`
}

type successRes struct {
	Expr   string  `json:"expr"`
	Answer float64 `json:"answer"`
}

type errorRes struct {
	Message string `json:"message"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("html/")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/calc/", calc)
	fmt.Println("listen and serve... http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func calc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var reqBody reqBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		res, _ := json.Marshal(errorRes{"invalid request"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	// parse
	expr, err := eval.Parse(reqBody.Expr)
	if err != nil {
		res, _ := json.Marshal(errorRes{fmt.Sprintf("failed to parse the input: %v", err)})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	// error check
	vars := make(map[eval.Var]bool)
	if err = expr.Check(vars); err != nil {
		res, _ := json.Marshal(errorRes{fmt.Sprintf("invalid expr: %v", err)})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	// Web電卓アプリでは変数利用は想定していない
	if len(vars) > 0 {
		res, _ := json.Marshal(errorRes{"variables are not supported"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	env := eval.Env{}
	res, _ := json.Marshal(successRes{reqBody.Expr, expr.Eval(env)})
	w.Write(res)
}
