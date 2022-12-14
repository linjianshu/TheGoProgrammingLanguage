package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
