package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var indexPage = template.Must(template.ParseFiles("tmpl/index.tmpl"))

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		if err := indexPage.Execute(w, "test"); err != nil {
			log.Fatal(err)
		}
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", handler)
	fmt.Println("listen and serve... http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
