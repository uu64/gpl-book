package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	noSuchItem         = "no such item"
	queryRequired      = "item and price are required"
	itemAlreadyExists  = "item already exists"
	invalidPriceFormat = "invalid price format"
)

var itemList = template.Must(template.ParseFiles("list.tmpl"))

var mu sync.Mutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	mu.Lock()
	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
	mu.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", queryRequired, item)
		return
	}

	mu.Lock()
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "%s: %q\n", noSuchItem, item)
	}
	mu.Unlock()
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", queryRequired, item)
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", itemAlreadyExists, item)
		return
	}

	dollar, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", invalidPriceFormat, item)
		return
	}

	mu.Lock()
	db[item] = dollars(dollar)
	mu.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", queryRequired, item)
		return
	}

	dollar, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", invalidPriceFormat, item)
		return
	}

	mu.Lock()
	if _, ok := db[item]; ok {
		db[item] = dollars(dollar)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "%s: %q\n", noSuchItem, item)
	}
	mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%s: %q\n", queryRequired, item)
		return
	}

	mu.Lock()
	if _, ok := db[item]; ok {
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "%s: %q\n", noSuchItem, item)
	}
	mu.Unlock()
}
