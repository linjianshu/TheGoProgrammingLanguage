package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Map(func(r rune) rune {
		return r + 1
	}, "aaabbbccc")
	fmt.Println(s)

	//闭包  函数内部对外部变量的引用  变量的生命周期并不一定由它的作用域决定
	f := square()    //初始化x
	fmt.Println(f()) //给x++ 计算得到1
	fmt.Println(f()) //给x++ 计算得到4
	fmt.Println(f()) //给x++ 计算得到9

}

//返回值是函数类型
func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
