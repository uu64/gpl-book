package main

import (
	"fmt"
	"sort"

	"github.com/uu64/gpl-book/ch07/ex08/track"
)

type byArtist []*track.Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*track.Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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

func main() {
	fmt.Println("default")
	track.PrintTracks(tracks)
	fmt.Println()

	fmt.Println("multi-tier sort (Year > Artist)")
	mts := track.MultiTierSort{
		Records:      tracks,
		PrimaryKey:   "Year",
		SecondaryKey: "Artist",
	}
	sort.Sort(mts)
	track.PrintTracks(mts.Records)
	fmt.Println()

	fmt.Println("sort.Stable (Year > Artist)")
	sort.Stable(byArtist(tracks))
	sort.Stable(byYear(tracks))
	track.PrintTracks(tracks)
	fmt.Println()
}
