package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	//直接使用sort.sort不行 因为没有实现len less swap方法 intsaresort内置了intslice 直接强转过去 而intslice已经实现了三个方法
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}
