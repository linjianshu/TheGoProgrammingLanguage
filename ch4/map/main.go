package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"ljs", "apple", "banana"}
	m := make(map[string]int)
	m["ljs"] = 24
	m["apple"] = 1
	m["banana"] = 2

	// map的遍历是无序的 所以只能对键值进行排序 然后按顺序获取值
	sort.Strings(names)
	for _, name := range names {
		fmt.Println("name: ", name, " age: ", m[name])
	}

	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))

}

//map不可比较 比较的话要自己写一个循环
func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok {
			return false
		} else if v1 != v2 {
			return false
		}
	}

	return true
}
