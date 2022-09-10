package main

import "fmt"

func main() {
	_, err := Parse(1)
	if err != nil {
		fmt.Println("like something wrong ... details: ", err.Error())
		return
	}

}

// Parse 假设这是个解析器 解析语言的语法用的
//难免这个解析器有考虑不周的地方 那么这个地方就可能存在bug
func Parse(intput interface{}) (s string, err error) {

	defer func() {
		//从特定的panic中恢复过来 防止程序宕机
		if p := recover(); p != nil {
			//errorf 不是打印输出 而是赋值的
			err = fmt.Errorf("internal error : %v", p)
		}
	}()
	s = intput.(string)
	return s, err
}
