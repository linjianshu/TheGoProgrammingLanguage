package main

import (
	"io"
	"time"
)

func main() {

}

// Artifact 共有特性接口
type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
	Resolution() (x, y int)
}

// Streamer 抽取共有特性 不用修改原接口结构
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Album struct {
}
type Book struct {
}
type Movie struct {
}
type Magazine struct {
}
type Podcast struct {
}
type TVEpisode struct {
}
type Track struct {
}
