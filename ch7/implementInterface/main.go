package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	//接口赋值之后 就从抽象类型变成具体类型了
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	//w = time.Second  编译错误 time.Second类型缺少 write方法

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	//rwc = w 编译错误 io.Writer缺少Read 和 Close方法
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", rwc)

	//_ = Length(9).String()  字面量无法取指针 因此编译错误
	//这个就可以 因为已经有内存地址了 调用会隐式完成取地址操作
	l := Length(9)
	l.String()

	_ = fmt.Stringer(&l)
	//_ = fmt.Stringer(l)  编译错误

	os.Stdout.Write([]byte("hello world"))
	os.Stdout.Close()

	fmt.Fprintln(os.Stdout, "hello world")

	//只有通过接口暴露的方法才可以调用 类型的其他方法则无法通过接口调用
	var w1 io.Writer
	w1 = os.Stdout
	w1.Write([]byte("hello world"))
	//w.Close() 编译错误 这个接口没有这个方法  尽管os.Stdout有Close方法
}

type Length int

func (l *Length) String() string {
	return fmt.Sprintf("%d cm.", l)
}
