package main

import "fmt"

func main() {
	// Incorrect()
	Correct()
}

func Incorrect() {
	funcList := make([]func(), 0)

	m := make(map[int]int, 100)

	for i := 0; i < 100; i++ {
		m[i] = i
	}

	//陷阱 捕获迭代变量
	for i := 0; i < 100; i++ {
		f := func() {
			//i是在不断变化的 f函数记录的是删除i地址所在的值 而i地址不变 值从0-99 最后终止于100
			delete(m, i)
		}
		funcList = append(funcList, f)
	}

	for _, f := range funcList {
		//因此这里面的f 都是删除m中key=100的函数 所以一直删不掉啊啊
		f()
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%v : %v\n", i, m[i])
	}
}

func Correct() {
	funcList := make([]func(), 0)

	m := make(map[int]int, 100)

	for i := 0; i < 100; i++ {
		m[i] = i
	}

	//陷阱 捕获迭代变量
	for i := 0; i < 100; i++ {
		//正确的做法是 使用一个局部变量 把迭代变量的值先拷贝过来
		j := i
		f := func() {
			delete(m, j)
		}
		funcList = append(funcList, f)
	}

	for _, f := range funcList {
		//因此这里面的f 都是删除m中key=j的函数
		f()
	}

	for i := 0; i < 100; i++ {
		fmt.Printf("%v : %v\n", i, m[i])
	}
}
