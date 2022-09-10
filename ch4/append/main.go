package main

import "fmt"

func main() {
	s := "hello world"
	slice := []byte(s)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &slice[0])

	slice = append(slice, " hello"...)
	fmt.Println(string(slice))

	//原来的数组不变
	fmt.Println(s)

	slice[5] = ','
	fmt.Println(string(slice))
	fmt.Println(s)

}
