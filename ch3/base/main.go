package main

import "fmt"

func main() {
	//八进制
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) //[1]副词告知重复使用第一个操作数
	//十六进制
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x) //#告知输出相应的前缀
}
