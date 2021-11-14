package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"

	"github.com/uu64/gpl-book/ch07/ex08/track"
)

var trackList = template.Must(template.ParseFiles("tracks.tmpl"))

var tracks = []*track.Track{
	track.New("Go", "Delilah", "From the Roots Up", 2012, track.Length("3m38s")),
	track.New("Go", "Moby", "Moby", 1992, track.Length("3m37s")),
	track.New("Go Ahead", "Alicia Keys", "As I Am", 2007, track.Length("4m36s")),
	track.New("Time Machine", "Alicia Keys", "ALICIA", 2020, track.Length("4m26s")),
	track.New("If I Ain't Got You", "Alicia Keys", "The Diary of Alicia Keys", 2003, track.Length("3m48s")),
	track.New("Ready 2 Go", "Martin Solveig", "Smash", 2011, track.Length("4m24s")),
	track.New("Ready 2 Go", "Martin Solveig", "Smash", 2011, track.Length("4m24s")),
	track.New("Under the Bridge", "Red Hot Chili Peppers", "Blood Sugar Sex Magik", 1992, track.Length("4m24s")),
}

var mts = track.MultiTierSort{
	Records: tracks,
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		primaryKey := parseQuery(r.URL)
		mts.PrimaryKey = primaryKey
		sort.Sort(mts)
		if err := trackList.Execute(w, mts); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/tracks/", handler)
	fmt.Println("listen and serve... http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func parseQuery(url *url.URL) (primaryKey string) {
	q := url.Query()
	primaryKey = q.Get("primaryKey")
	return
}
