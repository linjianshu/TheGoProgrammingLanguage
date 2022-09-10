package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//重复count次
	repeat := strings.Repeat("hello hello world world world", 2)
	fmt.Println(repeat)

	//中间的不会trim
	trim := strings.Trim(". hello , world .", ".")
	fmt.Println(trim)

	//字典序比较  z>a 越长越好
	compare := strings.Compare("hello", "hello")
	fmt.Println(compare)

	compare = strings.Compare("hello", "world")
	fmt.Println(compare)

	compare = strings.Compare("a", "z")
	fmt.Println(compare)

	compare = strings.Compare("a", "aa")
	fmt.Println(compare)

	replace := strings.Replace("hello , world world", "world", "ljs", 2)
	fmt.Println(replace)

	replaceAll := strings.ReplaceAll("hello , world world", "world", "ljs")
	fmt.Println(replaceAll)

	contains := strings.Contains("hello , world", "wor")
	fmt.Println(contains)

	containsRune := strings.ContainsRune("hello , 世界", '世')
	fmt.Println(containsRune)

	//一个或多个连续的空白
	fields := strings.Fields("   hel  lo    world   ")
	fmt.Println(fields)

	//bytes 包和strings方法类似
	i := bytes.Compare([]byte("a"), []byte("b"))
	fmt.Println(i)

	join := bytes.Join([][]byte{[]byte("你"), []byte("好")}, []byte(","))
	fmt.Println(string(join))

	//类似 fmt %q增加引号
	quote := strconv.Quote("hello")
	fmt.Println(quote)

	//去除引号
	unquote, _ := strconv.Unquote(quote)
	fmt.Println(unquote)

	//unicode包提供 是否大小写 是否十进制 是否字母数字
	letter := unicode.IsLetter(65)
	fmt.Println(letter)

	number := unicode.IsNumber(65)
	fmt.Println(number)

}
