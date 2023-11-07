package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := 4, 0
	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err.Error())
	return
	}
	fmt.Println(a, "除以", b, "的结果是 ", res)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("cannot be zero")
	}
	return a/b, nil
}