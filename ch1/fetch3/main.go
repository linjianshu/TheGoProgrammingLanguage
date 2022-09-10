package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	totalStart := time.Now()
	c := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, c)
	}
	for range os.Args[1:] {
		fmt.Println(<-c)
	}
	fmt.Println(time.Since(totalStart).Milliseconds())
}

func fetch(url string, c chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("get %s error , %v", url, err.Error())
		return
	}
	written, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		c <- fmt.Sprintf("copy %s error , %v", url, err.Error())
		return
	}

	c <- fmt.Sprintf("succeeded %s, copy %d bytes, cost %d Milliseconds.", url, written, time.Since(start).Milliseconds())
}
