package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

//输入的时候要包含单位 s ns ms
func main() {
	flag.Parse()
	//time.Duration 方法实现了接口fmt.Stringer
	fmt.Printf("sleeping for %v ...\n", period)
	time.Sleep(*period)
	fmt.Println("ended !!!")
}
