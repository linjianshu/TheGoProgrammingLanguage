package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	get, err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(all))
}
