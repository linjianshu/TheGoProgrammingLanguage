package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if m[scanner.Text()] {
			continue
		}
		fmt.Println(scanner.Text())
		m[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
