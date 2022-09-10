package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	//defer 是倒序输出的 panic会宕机 x==0的时候 输出堆栈信息
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
