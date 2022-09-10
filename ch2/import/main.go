package main

import (
	"bufio"
	"fmt"
	"os"
	"src/TheGoProgrammingLanguage/ch2/tempconv"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			float, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			//重写string方法
			fmt.Println(tempconv.Celsius(float), tempconv.Fahrenheit(float))
		}
	} else {
		for _, f := range os.Args[1:] {
			float, err := strconv.ParseFloat(f, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf :%v\n", err)
				return
			}
			celsius := tempconv.Celsius(float)
			fahrenheit := tempconv.Fahrenheit(float)
			fmt.Println(celsius, fahrenheit)
		}
	}
}
