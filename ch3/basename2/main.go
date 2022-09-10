package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/a/b.c.go"
	//如果不存在返回-1
	index := strings.LastIndex(s, "/")
	//-1+1刚好等于0 所以不会报错 设计巧妙
	s = s[index+1:]
	lastIndex := strings.LastIndex(s, ".")
	if lastIndex >= 0 {
		s = s[:lastIndex]
	}
	fmt.Println(s)

}
