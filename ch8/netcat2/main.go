package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dial.Close()
	go mustCopy(os.Stdout, dial)
	mustCopy(dial, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
