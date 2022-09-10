package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	bytes := sha256.Sum256([]byte("x"))
	sum256 := sha256.Sum256([]byte("X"))
	fmt.Printf("%X\n", bytes)
	fmt.Printf("%X\n", sum256)
	fmt.Println(bytes == sum256)
	fmt.Printf("%b\n", 'x')
	fmt.Printf("%b\n", 'X')

	fmt.Println(bytes)
	zero(&bytes)
	fmt.Println(bytes)
}

//数组在go中是值传递 传递后就算改变也是改变副本的值
//因此适合使用指针传递 虽然拷贝的是指针的值 但是指针指向的地址确实同一个地方
//因此能够修改原数组的值
func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
