package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	buffer := bytes.Buffer{}
	count := 0
	for {
		readRune, size, err := reader.ReadRune()
		if err == io.EOF {
			fmt.Println("EOF ", "read ", count, " bytes ...")
			fmt.Println(buffer.String())
			break
		} else if err != nil {
			fmt.Println(err)
			fmt.Println("read ", size, " bytes ...")
			break
		}
		buffer.WriteRune(readRune)
		count += size
	}
}
