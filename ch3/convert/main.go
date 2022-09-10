package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换出错
	f := 1e100
	i := int(f)
	fmt.Println(i)

	x := 123
	fmt.Printf("%T %[1]v\n", fmt.Sprintf("%d", x))
	y := strconv.Itoa(x)
	fmt.Printf("%T %[1]v\n", y)

	//转换具体的进制
	fmt.Println(strconv.FormatInt(123, 2))

	atoi, _ := strconv.Atoi(y)
	fmt.Printf("%T %[1]v\n", atoi)

	//转换成几进制 并用什么类型装 32?64
	parseInt, _ := strconv.ParseInt("123", 8, 64)
	fmt.Println(parseInt)
}
