package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	c := Celsius(26)
	fmt.Println(c.String())
	fmt.Println(c)        //不调用字符串也可以
	fmt.Printf("%v\n", c) //不调用字符串也可以
	fmt.Printf("%s\n", c) //不调用字符串也可以
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}

//重写string 方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}
