package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf *bytes.Buffer
	fmt.Printf("%p\n", buf)
	if debug {
		buf = new(bytes.Buffer)
		fmt.Printf("%p\n", buf)
	}

	f(buf)
	fmt.Println(buf.String())
}

func f(writer io.Writer) {
	//debug 为false的时候 接口中的动态类型是有的 只是动态值是nil 动态类型非空就可以判断接口非空 所以会进去
	if writer != nil {
		writer.Write([]byte("done !\n"))
	}
}

const debug = true
