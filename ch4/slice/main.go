package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	//虽说修改的是切片 但是切片终究还是指向数组的 也就是说底层的数组被修改了
	reverse(a[:])
	fmt.Println(a)
	fmt.Println(a[:])

	reset(&a)
	//把前n个左移 只需要简单的进行三次反转就可以了
	n := 2
	reverse(a[:n])
	reverse(a[n:])
	reverse(a[:])
	fmt.Println(a)
	fmt.Println(a[:])

	//slice 不能进行比较运算
	h1 := []byte("hello world")
	h2 := []byte("hello world")
	//bytes包 equal 为[]byte 比较提供函数
	equal := bytes.Equal(h1, h2)
	fmt.Println(equal)

	//其他的只能自己写函数来比较
	equalFunc := func(a []byte, b []byte) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	fmt.Println(equalFunc(h1, h2))

	//匿名函数 即写即用
	b := func(a []byte, b []byte) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}(h1, h2)
	fmt.Println(b)

}

//可以选择传递数组的指针 达到通过地址寻值 继而修改值的目的
func reset(prt *[6]int) {
	*prt = [...]int{0, 1, 2, 3, 4, 5}
}

//传递切片而不是数组 因为数组是值传递 而切片是包含指针、容量、长度的轻量数据类型 属于引用传递 或者说传递的是地址的副本
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
