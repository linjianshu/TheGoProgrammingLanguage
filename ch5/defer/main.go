package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//bigSlowOperation()

	result := double(4)
	fmt.Println(result)
	result = triple(4)
	fmt.Println(result)

}

//进来之后 先走trace方法 走到return 暂时不走留着
//跳过defer 继续执行
//执行完 走上次停着的
//延迟执行的函数 在return语句之后执行 并且可以更新函数的结果变量
//因为函数可以得到其外层函数作用域内的变量 所以延迟执行的匿名函数可以修改观察到的函数的返回结果
func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//这里会先执行x+x 然后显示hello world double(4)=(8) 然后结束
//注意这里的差别 实际上 trace返回值是func()
//这里恰恰defer的就是func() 保持一致性的看 trace必须要执行完才能得到一个返回值为func()的函数 以供最后return之后的defer 有函数执行
func double(x int) (result int) {
	defer func() {
		fmt.Println("hello world")
		fmt.Printf("double (%d)=(%d)\n", x, result)
	}()
	return x + x
}

func triple(x int) (result int) {
	//闭包了 明明是return之后执行的defer 按理说result和x已经释放 但是由于还保留着引用 因此不会被释放
	//同时还保持对结果值的修改
	defer func() {
		result = result + x
	}()
	return x + x
}
