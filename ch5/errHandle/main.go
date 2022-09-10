package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := WaitForServer("http://localhost:8088/GetPage")
	if err != nil {
		fmt.Fprintf(os.Stderr, "site is donw %v\n", err)
		create, _ := os.Create("log")
		logger := log.New(create, "wait: ", 0)
		logger.Println("site is down ", err)
		log.Fatalf("site is donw %v\n", err)
		os.Exit(1)
	}
}

func WaitForServer(url string) error {
	now := time.Now()
	for tries := 0; time.Since(now) < time.Minute; tries++ {
		_, err := http.Head(url)
		if err != nil {
			log.Printf("server not responding (%s) ; retrying ...", err)
			//指数退避策略
			log.Printf("go to sleep about %v seconds ...", time.Second<<tries)
			time.Sleep(time.Second << tries)
		} else {
			return nil
		}
	}
	return fmt.Errorf("server %s failed to respond after %s", url, time.Minute)
}
