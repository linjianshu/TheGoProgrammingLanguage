package main

import "fmt"

func main() {
	fmt.Println(nonEmpty([]string{"he", " ", "he"}))
	fmt.Println(nonEmpty2([]string{"he", " ", "he"}))
}

func nonEmpty(strings []string) []string {
	//原地 trim
	i := 0
	for _, s := range strings {
		if s != " " {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonEmpty2(strings []string) []string {
	//原地 trim
	//引用原始slice的新的零长度的slice
	s := strings[:0]
	for _, str := range strings {
		if str != " " {
			s = append(s, str)
		}
	}
	return s
}
