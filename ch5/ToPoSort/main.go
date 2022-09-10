package main

import (
	"fmt"
	"sort"
)

func main() {
	//拓扑排序
	var prereqs = map[string][]string{
		"algorithms":            {"data structures"},
		"calculus":              {"linear algebra"},
		"compilers":             {"data structures", "formal lauguages", "computer organization"},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal lauguages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	keys := make([]string, 0, len(prereqs))
	for key := range prereqs {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	res := make([]string, 0)

	visited := make(map[string]bool)

	for _, course := range keys {
		s := TopoSort(course, prereqs[course], prereqs, visited)
		res = append(res, s...)
	}

	for i, course := range res {
		fmt.Printf("%v:\t %s\n", i+1, course)
	}
}

func TopoSort(course string, m []string, prereqs map[string][]string, visited map[string]bool) []string {
	if visited[course] {
		return []string{}
	}

	ans := make([]string, 0)
	for _, v := range m {
		latterCourse := TopoSort(v, prereqs[v], prereqs, visited)
		ans = append(ans, latterCourse...)
	}
	ans = append(ans, course)
	visited[course] = true
	return ans
}
