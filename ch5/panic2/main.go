package main

import (
	"fmt"
	"runtime"
)

func main() {
	//defer语句放置的位置很考究 如果放在可能出现问题的位置之前 才会在问题错误的时候 return回来 继续执行defer
	defer printStack()
	f(3)
	//defer语句放置的位置很考究 如果放在可能出现问题的位置之后 在问题错误出现的时候 就直接从错误的地方断了
	//因此如果放在f(3)的后边 压根不会走到里面去
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("%v\n", string(buf[:n]))
}

func f(x int) {
	//defer 是倒序输出的 panic会宕机 x==0的时候 输出堆栈信息
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
