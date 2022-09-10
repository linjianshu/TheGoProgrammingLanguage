package main

import "fmt"

func main() {

	type Currency int
	const (
		USD Currency = iota
		RMB
	)
	c := [...]rune{USD: '$', RMB: '￥'}
	fmt.Printf("%T %[1]v %[1]c\n", c)

}
