package main

import "fmt"

var m = make(map[string]int)

func main() {
	//有时候 键要使用切片类型 那么只能使用一个帮助函数k来转换
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func add(list []string) {
	m[k(list)]++
}

func count(list []string) int {
	return m[k(list)]
}
