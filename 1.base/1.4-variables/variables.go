package main

import (
	"fmt"
	"strconv"
)

const (
	f = iota
	g
)

func main() {
	var a string
	a = "1"
	var b = "2"
	var c, d, e = 3, 4, 5
	fmt.Printf("a%s b%s c%d d%d e%d f%d g%d \n",a,b,c,d,e,f,g)
	fmt.Println("==============")
	var aInt int = 17
	// 一般用这种方式强制转
	fmt.Printf("转float64 %f  \n", float64(aInt))
	fmt.Printf("转string %v  \n", strconv.Itoa(aInt))
	fmt.Printf("转float64 %f  \n", float64(aInt))
}