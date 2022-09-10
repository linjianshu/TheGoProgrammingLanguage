package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	b := flag.Bool("n", false, "omit trailing newline")
	s := flag.String("s", " ", "separator")

	//很关键
	flag.Parse()

	fmt.Println(strings.Join(flag.Args(), *s))
	if !*b {
		fmt.Println()
	}

}
