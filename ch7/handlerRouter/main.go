package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	http.ListenAndServe(":8080", db)
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//缺点就是 逻辑都在一块 看的像屎山
	switch r.URL.Path {
	case "/list":
		for key, value := range d {
			fmt.Fprintf(w, "%s: %s\n", key, value)
		}
	case "/price":
		itemName := r.URL.Query().Get("item")
		if v, ok := d[itemName]; ok {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s: %s\n", itemName, v)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "the thing [%s] that you search is no found\n", itemName)
		}
	default:
		//http.Error工具使用同样达到效果
		http.Error(w, fmt.Sprintf("no such page: %s\n", r.URL), http.StatusNotFound)
	}

}
