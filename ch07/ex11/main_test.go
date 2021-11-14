package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var db database

type item struct {
	k string
	v dollars
}

func initialize(items ...item) {
	db = make(map[string]dollars)
	for _, item := range items {
		db[item.k] = dollars(item.v)
	}
}

func TestCreate(t *testing.T) {
	initialize(item{"shoes", 50}, item{"socks", 5})
	tests := []struct {
		url     string
		code    int
		message string
		item    item
	}{
		{"http://localhost:8080/create", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/create?item=hat", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/create?price=100", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/create?item=hat&price=aaa", http.StatusBadRequest, invalidPriceFormat, item{"", 0}},
		{"http://localhost:8080/create?item=shoes&price=200", http.StatusBadRequest, itemAlreadyExists, item{"", 0}},
		{"http://localhost:8080/create?item=hat&price=100", http.StatusOK, "", item{"hat", 100}},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()

		db.create(w, req)
		resp := w.Result()
		if resp.StatusCode != test.code {
			t.Errorf("GET %s returns code: %v\n", test.url, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.HasPrefix(string(body), test.message) {
			t.Errorf("GET %s returns body: %v\n", test.url, string(body))
		}

		if test.code == http.StatusOK {
			if price, ok := db[test.item.k]; !ok || price != test.item.v {
				t.Errorf("missing items %v: %v\n", test.item, db)
			}
		}
	}
}

func TestUpdate(t *testing.T) {
	initialize(item{"shoes", 50}, item{"socks", 5})
	tests := []struct {
		url     string
		code    int
		message string
		item    item
	}{
		{"http://localhost:8080/update", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/update?item=shoes", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/update?price=100", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/update?item=shoes&price=aaa", http.StatusBadRequest, invalidPriceFormat, item{"", 0}},
		{"http://localhost:8080/update?item=hat&price=100", http.StatusNotFound, noSuchItem, item{"", 0}},
		{"http://localhost:8080/update?item=shoes&price=100", http.StatusOK, "", item{"shoes", 100}},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()

		db.update(w, req)
		resp := w.Result()
		if resp.StatusCode != test.code {
			t.Errorf("GET %s returns code: %v\n", test.url, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.HasPrefix(string(body), test.message) {
			t.Errorf("GET %s returns body: %v\n", test.url, string(body))
		}

		if test.code == http.StatusOK {
			if price, ok := db[test.item.k]; !ok || price != test.item.v {
				t.Errorf("missing items %v: %v\n", test.item, db)
			}
		}
	}
}

func TestDelete(t *testing.T) {
	initialize(item{"shoes", 50}, item{"socks", 5})
	tests := []struct {
		url     string
		code    int
		message string
	}{
		{"http://localhost:8080/delete", http.StatusBadRequest, queryRequired},
		{"http://localhost:8080/delete?item=hat", http.StatusNotFound, noSuchItem},
		{"http://localhost:8080/delete?item=shoes", http.StatusOK, ""},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()

		db.delete(w, req)
		resp := w.Result()
		if resp.StatusCode != test.code {
			t.Errorf("GET %s returns code: %v\n", test.url, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.HasPrefix(string(body), test.message) {
			t.Errorf("GET %s returns body: %v\n", test.url, string(body))
		}

		key := req.URL.Query().Get("item")
		if test.code == http.StatusOK {
			if _, ok := db[key]; ok {
				t.Errorf("failed to delete %s: %v\n", key, db)
			}
		}
	}
}

func TestPrice(t *testing.T) {
	initialize(item{"shoes", 50}, item{"socks", 5})
	tests := []struct {
		url     string
		code    int
		message string
		item    item
	}{
		{"http://localhost:8080/price", http.StatusBadRequest, queryRequired, item{"", 0}},
		{"http://localhost:8080/price?item=hat", http.StatusNotFound, noSuchItem, item{"", 0}},
		{"http://localhost:8080/price?item=shoes", http.StatusOK, "", item{"shoes", 50}},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		w := httptest.NewRecorder()

		db.price(w, req)
		resp := w.Result()
		if resp.StatusCode != test.code {
			t.Errorf("GET %s returns code: %v\n", test.url, resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)

		if test.code == http.StatusOK {
			if string(body) != fmt.Sprintf("%s\n", test.item.v) {
				t.Errorf("GET %s returns body: %v\n", test.url, string(body))
			}
		} else {
			if !strings.HasPrefix(string(body), test.message) {
				t.Errorf("GET %s returns body: %v\n", test.url, string(body))
			}
		}
	}
}
