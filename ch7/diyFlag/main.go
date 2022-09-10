package main

import (
	"flag"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%v°\n", float64(c))
}

func (c *Celsius) Set(s string) error {
	//单位
	var unit string
	//值
	var value float64
	//从字符串中解析
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "", "C", "°C":
		*c = Celsius(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q\n", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := value
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
