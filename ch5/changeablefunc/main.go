package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum(0))
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	fmt.Printf("%T\n", sum)
	fmt.Printf("%T\n", g)

	errorf(24, "%s %s\n", "hello ", "world")
}

func sum(ints ...int) int {
	var sum int
	for _, v := range ints {
		sum += v
	}
	return sum
}

func g([]int) int {
	return 0
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %v\t", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
