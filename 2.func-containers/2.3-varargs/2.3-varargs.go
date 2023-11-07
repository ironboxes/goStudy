package main

import "fmt"

func main() {
	// fmt.Println(test(1,2,3,4,5))
	fmt.Println(test1(1,2,true, "2",3.1,4.2))
}

func test(t ...int) (res int) {
	for _, v := range t {
		res += v
	}
	return
}


func test1(t ...interface{}) (res int) {
	for _, v := range t {
		switch val:= v.(type) {
		case int:
			res += val
		case float32:
			res += int(val)
		case float64:
			res += int(val)
		}
	}
	return
}


