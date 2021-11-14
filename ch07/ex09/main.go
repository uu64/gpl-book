package main

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/uu64/gpl-book/ch07/ex08/track"
)

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
	router := gin.Default()
	router.LoadHTMLFiles("tracks.tmpl")
	router.GET("/tracks", func(c *gin.Context) {
		primaryKey := c.DefaultQuery("primaryKey", "")
		mts.PrimaryKey = primaryKey
		sort.Sort(mts)
		c.HTML(http.StatusOK, "tracks.tmpl", gin.H{
			"TotalCount": len(mts.Records),
			"Tracks":     mts.Records,
			"Key":        mts.PrimaryKey,
		})
	})
	router.Run(":8080")
}
