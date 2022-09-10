package main

import (
	"bytes"
	"fmt"
)

func main() {
	toString := intsToString([]int{1, 2, 3})
	fmt.Println(toString)
}

func intsToString(values []int) string {
	//buffer高效
	buffer := bytes.Buffer{}
	buffer.WriteByte('[')
	for i, value := range values {
		if i > 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%d", value)
	}
	buffer.WriteByte(']')
	return buffer.String()
}
