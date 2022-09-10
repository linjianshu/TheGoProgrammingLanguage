package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

func (p Path) Distance() float64 {
	var sum float64
	for i, point := range p {
		if i > 0 {
			sum += point.Distance(p[i-1])
		}
	}
	return sum
}

func main() {
	p := Point{
		X: 1,
		Y: 1,
	}
	q := Point{
		X: 5,
		Y: 1,
	}
	distance := p.Distance(q)
	fmt.Println(distance)

	path := Path{p, q, Point{
		X: 5,
		Y: 4,
	}, Point{
		X: 1,
		Y: 1,
	}}
	f := path.Distance()
	fmt.Println(f)

}
