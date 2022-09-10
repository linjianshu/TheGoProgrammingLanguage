package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//不指定类型 精度非常高
	const Ipv4Len = 4

	//同时指定类型和值 但是精度就锁定在类型上了
	const noDelay time.Duration = 0

	//和类型进行计算 锁定了类型 锁定了精度
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", Ipv4Len)
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)

	const (
		a = 1
		b //还是1
		c = 2
		d //还是2
	)

	fmt.Printf("%T %[1]v\n", a)
	fmt.Printf("%T %[1]v\n", b)
	fmt.Printf("%T %[1]v\n", c)
	fmt.Printf("%T %[1]v\n", d)

	type weekday int
	const (
		mon weekday = iota //iota=0
		tus                //递增
		wens
		turs
		frid
		satd
		sund
	)

	fmt.Printf("%T %[1]v\n", mon)
	fmt.Printf("%T %[1]v\n", tus)
	fmt.Printf("%T %[1]v\n", wens)
	fmt.Printf("%T %[1]v\n", turs)
	fmt.Printf("%T %[1]v\n", frid)
	fmt.Printf("%T %[1]v\n", satd)
	fmt.Printf("%T %[1]v\n", sund)

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)

	fmt.Printf("%T %[1]v\n", kb)
	fmt.Printf("%T %[1]v\n", mb)
	fmt.Printf("%T %[1]v\n", gb)
	fmt.Printf("%T %[1]v\n", tb)
	fmt.Printf("%T %[1]v\n", pb)
	fmt.Printf("%T %[1]v\n", eb)
	//fmt.Printf("%T %[1]v\n", zb)  //打印不出来 没有容器类型承接的下
	//fmt.Printf("%T %[1]v\n", yb)

	fmt.Println(eb / pb) //两个无法表示的大常量却可以相除得到合法的量

	//正是因为pi常量没有锁定类型 所以类型转换的时候不会丢失精度
	var x float32 = math.Pi
	var y = math.Pi
	var z complex128 = math.Pi
	fmt.Printf("%T %[1]v\n", x)
	fmt.Printf("%T %[1]v\n", y)
	fmt.Printf("%T %[1]v\n", z)

	var v = FlagMulticast | FlagUp
	fmt.Printf("%08b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%08b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%08b %t\n", v, IsUp(v))
	fmt.Printf("%08b %t\n", v, IsCast(v))

	fmt.Printf("%T %[1]v\n", 0)       //无类型整数
	fmt.Printf("%T %[1]v\n", 0.0)     //无类型浮点数
	fmt.Printf("%T %[1]v\n", "hello") //无类型字符串
	fmt.Printf("%T %[1]v\n", 'h')     //无类型字符
	fmt.Printf("%T %[1]v\n", 0i)      //无类型负数
	fmt.Printf("%T %[1]v\n", true)    //无类型布尔
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota //向上  				00000001
	FlagBroadcast                      //支持广播访问 		00000010
	FlagLoopback                       //是环回接口  			00000100
	FlagPointToPoint                   //属于点对点链路   		00001000
	FlagMulticast                      //支持多路广播访问  	00010000
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	//两种应该都可以吧
	return (v&FlagBroadcast == FlagBroadcast) || (v&FlagMulticast == FlagMulticast)
	return v&(FlagBroadcast|FlagMulticast) != 0
}
