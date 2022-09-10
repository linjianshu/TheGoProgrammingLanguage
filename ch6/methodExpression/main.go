package main

import "math"

func main() {

}

type Path []Point

func (p Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		//方法表达式会将接受者作为第一个参数 所以op的类型和Point.Add的类型是一致的
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for _, point := range p {
		point = op(point, offset)
	}
}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}
func (p Point) Sub(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}
