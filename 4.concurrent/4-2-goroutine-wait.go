package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	//test()
	//test1()
	//test2()
	//errFunc()
	correctFunc()
}

func test() {
	go say("hello world")
  time.Sleep(time.Second*1)
  fmt.Println("over!")
}

func test1() {
	done := make(chan bool)
	go func ()  {
		for i := 0; i < 3; i++ {
			fmt.Println("hello world")
		}
		done <- true
	}()
	<-done
	fmt.Println("over!")
}

func test2() {
	var wg sync.WaitGroup
	wg.Add(2)

	go say2("hello world!",&wg)
	fmt.Println("over")
	go say2("hello world!",&wg)
	wg.Wait()
}

func say2(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func errFunc() {
	var wg sync.WaitGroup
	sList := []string{"a", "b"}
	wg.Add(len(sList))
	for _, d := range sList {
		go func() {
			defer wg.Done()
			fmt.Println(d)
		}()
	}
	wg.Wait()
}

func correctFunc() {
	var wg sync.WaitGroup
	sList := []string{"a", "b"}
	wg.Add(len(sList))
	for _ , d:= range sList {
		go func (s string)  {
			defer wg.Done()
			fmt.Println(s)
		}(d)
	}
	wg.Wait()
}
