package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonaccid(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		// \r是软回车 表示回到这一行的开始 达到覆盖的效果
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}
