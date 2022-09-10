package main

import (
	"bytes"
	"fmt"
)

func main() {
	//接口是一个类型
	var a interface{}

	a = "hello world"
	a = 1
	a = false
	a = 1.00
	a = map[string]string{
		"ljs": "dsb",
	}
	a = []int{1}
	a = new(bytes.Buffer)

	//一旦对接口类型赋值 尽管原具体类型有x种方法 也只能使用接口类型种定义的方法
	//如何取得原始值 就需要使用类型断言来做了
	buffer := a.(*bytes.Buffer)
	fmt.Println(buffer.String())
}
