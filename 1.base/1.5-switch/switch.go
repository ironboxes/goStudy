package main

import "fmt"

func main() {
	var a interface{}
	a = "1"
	switch t := a.(type) {
	case int:
		fmt.Println("int type",t)
	case string:
		fmt.Println("string type",t)
	default:
		fmt.Println("unkonw type",t)
	}
}