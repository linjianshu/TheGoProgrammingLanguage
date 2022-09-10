package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)    //集合1,5
	fmt.Printf("%08b\n", y)    //集合1,2
	fmt.Printf("%08b\n", x&y)  //交集1
	fmt.Printf("%08b\n", x|y)  //并集1,2,5
	fmt.Printf("%08b\n", x^y)  //异或
	fmt.Printf("%08b\n", x&^y) //与非门 先对y非 然后对x与

	for i := uint8(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			//判定元素
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1) //集合2,6
	fmt.Printf("%08b\n", x>>1) //集合0,4
}
