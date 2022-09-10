package main

import "fmt"

func main() {
	fmt.Println(remove([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(remove2([]int{1, 2, 3, 4, 5}, 2))
}

//从中删除一个元素 并保持顺序
func remove(slice []int, pos int) []int {
	//原地删除
	//把当前的dst 用后面的src填充
	//然后slice的长度-1就可以了
	copy(slice[pos:], slice[pos+1:])
	return slice[:len(slice)-1]
}

//从中删除一个元素 不保持顺序
func remove2(slice []int, pos int) []int {
	slice[pos] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
