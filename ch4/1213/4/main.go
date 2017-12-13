package main

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
	"sort"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track
type byLength []*Track

var (
	tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
)

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "length")
	fmt.Fprintf(tw, format, "----", "----", "----", "----", "----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()

}

func (artist byArtist) Len() int {
	return len(artist)
}

func (artist byArtist) Less(i, j int) bool {
	return artist[i].Artist < artist[j].Artist
}

func (artist byArtist) Swap(i, j int) {
	artist[i], artist[j] = artist[j], artist[i]
}

func (length byLength) Len() int {
	return len(length)
}

func (length byLength) Less(i, j int) bool {
	return length[i].Length > length[j].Length
}

func (length byLength) Swap(i, j int) {
	length[i], length[j] = length[j], length[i]
}

func main() {
	var b byArtist = tracks
	sort.Sort(b)
	printTracks(b)

	fmt.Println()
	var l byLength = tracks
	sort.Sort(l)
	printTracks(l)

}
