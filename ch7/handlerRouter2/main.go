package main

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	//mux := http.DefaultServeMux
	//mux.HandleFunc("/list", db.list)
	//mux.HandleFunc("/price", db.price)
	//http.ListenAndServe(":8080", mux)

	//等同于下面这个 更加简单的版本

	//提供了一个全局的serverMux实例 DefaultServeMux 以及包级别的注册函数 http.Handle 和 http.HandleFunc
	//要让defaultservemux作为主处理程序 只需要传给nil就可以了
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.ListenAndServe(":8080", nil)
}

type database map[string]dollars

func (d database) list(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	for key, value := range d {
		fmt.Fprintf(w, "%s: %s\n", key, value)
	}
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("item")
	if v, ok := d[itemName]; ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s: %s\n", itemName, v)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "the thing [%s] that you search is no found\n", itemName)
	}
}
