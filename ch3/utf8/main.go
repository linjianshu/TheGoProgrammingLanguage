package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		runeInString, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf(" rune : %c, size : %v\n", runeInString, size)
		i += size
	}

	for k, v := range s {
		fmt.Printf("i = %v, word = %q, rune = %[2]d\n", k, v)
	}

	count := 0
	for range s {
		count++
	}
	fmt.Println("utf-8 count: ", count)

	fmt.Println(string(65))
	fmt.Println(string(1234567))
}
