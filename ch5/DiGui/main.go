package main

import "fmt"

func main() {
	strings := make([]string, 2, 2)
	strings[0] = "a"
	fmt.Println(strings)
	Digui2(strings)
	fmt.Println(strings)
	fmt.Println(len(strings))
	fmt.Println(strings[:])
}

// Digui 改变 直接传递切片的指针地址 强行将切片的长度和容量增加了
//并且修改了切片指向的底层数组的值
//函数返回后 由于原切片的长度和容量都增加了 所以能看得到+
func Digui(strs *[]string) {
	*strs = append(*strs, "hello")

}

// Digui1 不改变 实际上改变了底层数组的值的 这是毋庸置疑的 但是
//切片是一种数据结构 切片是一种 包含 底层数组 长度 容量的数据结构
//传递进来之前 他的长度和容量都固定了 传进来的是切片的副本 也就是该数据结构的副本
//在这个副本中 长度增加了 底层数组的数据修改了
//但是函数返回后 原来的切片中 底层数组值确实变了 但是长度和容量没变
//同时无法修改已定义切片的长度和容量 因此不是值没变 而是看不到 到不了
func Digui1(strs []string) {
	strs = append(strs, "hello")

}

// Digui2 改变了的验证
//我们故意将strings的长度设置为2 在传入函数之前 切片的长度和容量都是2
//传进来前赋值了[0] 传进来后 拷贝了该切片的副本
//随后对副本指向的底层数组进行修改 数组中将[1]的值修改了
//函数返回后 由于原切片的长度和容量都为2 所以能看到[1]底层数组的值
func Digui2(strs []string) {
	strs[1] = "hello"
}
