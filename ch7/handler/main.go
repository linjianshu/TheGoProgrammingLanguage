package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	//handler 是一个接口 只要是实现了该接口的类型都能传递进来 显然db类型实现了serveHTTP方法
	log.Fatal(http.ListenAndServe("localhost:8080", db))

}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for key, value := range d {
		fmt.Fprintf(w, "%s: %s\n", key, value)
	}
}
