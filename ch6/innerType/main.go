package main

import "sync"

func main() {
	lookup("user")
	lookup1("user")
}

var (
	mu      = sync.Mutex{}
	mapping = make(map[string]string)
)

//普通方法
func lookup(key string) string {
	mu.Lock()
	s := mapping[key]
	mu.Unlock()
	return s
}

//由于sync.mutex是内嵌的 因此能够通过内嵌类型 直接使用lock和unlock方法 同时新的变量名更加贴切
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: map[string]string{},
}

func lookup1(key string) string {
	cache.Lock()
	s := cache.mapping[key]
	cache.Unlock()
	return s
}
