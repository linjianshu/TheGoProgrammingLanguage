package main

import (
	"log"
	"os"
)

func main() {

}

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("failed :%v\n", err)
	}
	log.Printf("Working directory = %s\n", cwd)
}
