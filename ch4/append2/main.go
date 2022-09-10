package main

import "fmt"

func main() {
	ints := make([]int, 3, 5)
	ints = []int{1, 2, 3}
	//索引超出len-1范围会报错
	//fmt.Println(ints[3])
	//fmt.Println(ints[4])
	ints = append(ints, 4)
	fmt.Println(len(ints), cap(ints))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//模拟append 功能  实现扩容
func appendInt(x []int, y int) []int {
	var z []int

	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		capZ := zlen * 2
		z = make([]int, zlen, capZ)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
