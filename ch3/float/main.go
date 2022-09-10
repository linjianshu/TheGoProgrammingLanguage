package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 16777216
	fmt.Println(f == f+1) //true
	const (
		Avogadro = 6.02214129e23  //科学计数法
		Planck   = 6.62606957e-34 //科学计数法
	)

	for i := 0; i < 11; i++ {
		//宽度超出8就超出了 没有超出就空格占位
		fmt.Printf("x =%d e^x = %8.3f\n", i, math.Exp(float64(i)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, 1/-z, z/z)

	fmt.Println(math.IsNaN(z / z))
	fmt.Println(math.NaN())
	naN := math.NaN()
	fmt.Println(naN, naN == naN, naN != naN, naN < naN, naN > naN)
}
