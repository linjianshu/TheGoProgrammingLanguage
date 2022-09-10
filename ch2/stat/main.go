package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	stat, err := file.Stat()
	file.Close()
	fmt.Println(stat)

}
