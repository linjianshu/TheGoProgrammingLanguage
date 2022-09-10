package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		readRune, size, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if readRune == unicode.ReplacementChar && size == 1 {
			invalid++
			continue
		}

		utflen[size]++
		counts[readRune]++
	}

	for r, i := range counts {
		fmt.Printf("%q\t%d\n", r, i)
	}

	for i, i2 := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, i2)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters \n", invalid)
	}
}
