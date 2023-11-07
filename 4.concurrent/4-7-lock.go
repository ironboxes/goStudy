package main

import (
	"fmt"
	"sync"
	"time"
)

var s []int
var lock sync.Mutex
func main() {
	//test()
	test1()
}

func appendValue(i int) {
	s = append(s, i)
}

func appendValue1(i int) {
	lock.Lock()
	s = append(s, i)
	lock.Unlock()
}

func test() {
	for i := 0; i < 10000; i++ {
		go appendValue(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println(len(s))
}

func test1() {
		for i := 0; i < 10000; i++ {
		go appendValue1(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println(len(s))
}