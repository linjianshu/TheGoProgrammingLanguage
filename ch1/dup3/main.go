package main

import (
	"fmt"
	"io/ioutil"
	"os"
	strings2 "strings"
)

func main() {
	m := make(map[string]int)
	strings := os.Args[1:]
	fmt.Println(strings)

	for _, s := range strings {
		file, err := ioutil.ReadFile(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		for _, word := range strings2.Split(string(file), "\n") {
			m[word]++
		}
	}

	for k, v := range m {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
