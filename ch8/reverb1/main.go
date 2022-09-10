package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		os.Exit(1)
		fmt.Println(err.Error())
		return
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			os.Exit(1)
			fmt.Println(err.Error())
			return
		}
		go handleConn(accept)
	}
}

func handleConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		echo(c, scanner.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
