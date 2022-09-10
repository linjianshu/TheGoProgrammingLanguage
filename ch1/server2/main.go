package main

import (
	"fmt"
	"net/http"
	"sync"
)

var m sync.Mutex
var counter int

func main() {
	http.HandleFunc("/", urlHandler)
	http.HandleFunc("/count", countHandler)
	http.ListenAndServe("localhost:8000", nil)

}

func countHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	fmt.Fprintf(w, "%v", counter)
	m.Unlock()
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	counter++
	m.Unlock()
	fmt.Fprintf(w, "%q", r.URL.Path)
}
