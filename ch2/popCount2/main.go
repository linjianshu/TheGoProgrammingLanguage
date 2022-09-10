package main

import "fmt"

func main() {
	count := popCount(4)
	fmt.Println(count)
}

func popCount(x int) int {
	count := 0
	for x != 0 {
		//b-k算法 和x-1进行与运算 直到x==0 运算的次数就是1的次数
		x = x & (x - 1)
		count++
	}
	return count
}
