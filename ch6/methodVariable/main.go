package main

import (
	"fmt"
	"math"
	"time"
)

//方法变量与表达式
func main() {
	//正常调用方法
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

	//方法变量
	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)
	distance = distanceFromP(q)
	fmt.Println(distance)

	//方法表达式  把接收者转换成第一个形参 好像静态方法
	distanceFunc := Point.Distance
	fmt.Printf("%T\n", distanceFunc)
	distance = distanceFunc(p, q)
	fmt.Println(distance)

	r := rocket{}
	time.AfterFunc(10*time.Second, func() { r.launch() })
	time.AfterFunc(10*time.Second, r.launch)

}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type rocket struct {
}

func (r rocket) launch() {

}
