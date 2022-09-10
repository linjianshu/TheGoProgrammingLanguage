package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))
}

func comma(s string) string {
	buffer := bytes.Buffer{}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if count == 3 {
			buffer.WriteByte(',')
			count = 0
		}
		buffer.WriteByte(s[i])
		count++
	}
	res := buffer.Bytes()
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return string(res)
}
