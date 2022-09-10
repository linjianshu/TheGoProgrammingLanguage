package main

import (
	"fmt"
)

type BytesCount int

func (b *BytesCount) Write(bytes []byte) (n int, err error) {
	s := string(bytes)
	*b += BytesCount(len(s))
	return len(s), nil
}
func main() {
	var a BytesCount = 1
	fprintf, _ := fmt.Fprintf(&a, "%s\n", "hello world")
	fmt.Println(fprintf)

}
