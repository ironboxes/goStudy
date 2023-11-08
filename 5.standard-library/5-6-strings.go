package main

import (
	"fmt"
	"strings"
)

func main() {
	test()
}

func test() {
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString(" ")
	builder.WriteString("world!")
	s := builder.String()
	fmt.Println(s)
}
