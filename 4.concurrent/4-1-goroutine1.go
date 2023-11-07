package main

import (
	"fmt"
	"time"
)

func main() {
	//test()
	test1()
}

func test() {
	go fun()
	fmt.Println("this is main")
	time.Sleep(time.Second)
}

func test1() {
	for i := 0; i < 10; i++ {
		go Add(i,i)
	}
}

func fun() {
	fmt.Println("you may can see this");
}

func Add(a, b int) {
	c := a + b
	fmt.Println(c)
}