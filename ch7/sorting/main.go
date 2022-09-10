package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

var tracks = []*Track{
	{
		Title:  "Go",
		Artist: "Delilah",
		Album:  "From the Roots Up",
		Year:   2012,
		Length: 1,
	}, {
		Title:  "Go",
		Artist: "Moby",
		Album:  "Moby",
		Year:   2007,
		Length: 1,
	}, {
		Title:  "Go Ahead",
		Artist: "Alicia Keys",
		Album:  "As I Am",
		Year:   2007,
		Length: 1,
	}, {
		Title:  "Ready 2 Go",
		Artist: "Martin Solveig",
		Album:  "Smash",
		Year:   2007,
		Length: 1,
	},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, track := range tracks {
		fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	tw.Flush()
}

func main() {
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	fmt.Fprintf(os.Stdout, "%v\n", "-------------------------------------------------------------")
	fmt.Println("reverse")
	fmt.Fprintf(os.Stdout, "%v\n", "-------------------------------------------------------------")
	//reverse 只是返回元数据  元数据中的Less方法的i和j调换了一下 以此打到颠倒的作用 最后还是需要调用的辣
	reverse := sort.Reverse(byArtist(tracks))
	sort.Sort(reverse)
	printTracks(tracks)

	sort.Sort(byYearThenArtist{tracks, func(i, j int) bool {
		//年份不等 先比年份
		if tracks[i].Year != tracks[j].Year {
			return tracks[i].Year < tracks[j].Year
		} else {
			//年份相等 就比艺术家
			return tracks[i].Artist < tracks[j].Artist
		}
	}})
	printTracks(tracks)
}

//自定义排序 按照艺术家来排序
type byArtist []*Track

func (b byArtist) Len() int {
	return len(b)
}

func (b byArtist) Less(i, j int) bool {
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//不一定只使用与slice 其他数据结构也可以的  自定义排序 先按照年份 在按照艺术家
type byYearThenArtist struct {
	t    []*Track
	less func(i, j int) bool
}

func (b byYearThenArtist) Len() int {
	return len(b.t)
}

func (b byYearThenArtist) Less(i, j int) bool {
	return b.less(i, j)
}

func (b byYearThenArtist) Swap(i, j int) {
	b.t[i], b.t[j] = b.t[j], b.t[i]
}
