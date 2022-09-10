package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func addRune(r rune) rune {
	return r + 1
}
func main() {
	s := square
	fmt.Println(s(3))

	s = negative
	fmt.Println(s(3))
	fmt.Printf("%T \n", s)

	//传递一个函数,让字符串变成另一个字符串
	s2 := strings.Map(addRune, "abcdef")
	fmt.Println(s2)
}
