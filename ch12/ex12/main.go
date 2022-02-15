package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/uu64/gpl-book/ch12/ex12/params"
)

// type Data struct {
// 	Labels     []string `http:"l"`
// 	MaxResults int      `http:"max"`
// 	Exact      int      `positive:"x"`
// }

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	type Data struct {
		Labels []string `http:"l"`
		Mail   string   `mail:"m"`
		Num    int      `positive:"n"`
	}
	var data Data
	// data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
	// d1 := &Data{
	// 	[]string{"golang", "programming"}, 150, true,
	// }
	// url, _ := params.Pack("http://localhost:12345/search", d1) // ignoring errors
	// fmt.Println(url)
}
