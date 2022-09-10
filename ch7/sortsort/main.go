package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"apple", "bananaA", "bananaB", "banana", "a"}
	sort.Strings(strings)
	fmt.Println(strings)

	slice := StringSlice{"apple", "bananaA", "bananaB", "banana", "a"}
	sort.Sort(slice)
	fmt.Println(slice)
}

type StringSlice []string

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
