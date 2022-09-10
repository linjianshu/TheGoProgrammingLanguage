package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]int)
	strings := os.Args[1:]
	if len(strings) == 0 {
		countLines(os.Stdin, m)
	} else {
		for _, s := range strings {
			file, err := os.Open(s)
			if err != nil {
				fmt.Println(os.Args[:])
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, m)
			file.Close()
		}
	}

	for k, v := range m {
		if v > 1 {
			fmt.Printf("%d\t %s\n", v, k)
		}
	}
}

func countLines(file *os.File, m map[string]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m[scanner.Text()]++
	}
}
