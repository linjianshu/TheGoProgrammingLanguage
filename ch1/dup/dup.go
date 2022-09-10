package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m[scanner.Text()]++
	}

	for k, v := range m {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
