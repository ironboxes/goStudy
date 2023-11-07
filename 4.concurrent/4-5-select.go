package main

import (
	"fmt"
	"time"
)

func main() {
	//test(99)
	//test1()
	//test2()
	A()
}

func test(num int) {
	switch {
	case num >= 90:
		fmt.Println("goood!")
	default:
		fmt.Println("bad!")
	}
}

func test1() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func ()  {
		select {
		case data, ok := <-chanInt:
			if ok {
				fmt.Println(data)
			}
		default:
			fmt.Println("全部阻塞")
		}
	}()
	time.Sleep(time.Second)
	chanInt <-1
}

func test2() {
	c := make(chan int)
	defer close(c)
	go func ()  {
		for {
			select {
			case i, ok:= <-c:
				if ok {
					fmt.Println(i)
				}
			default:
				fmt.Println("default")
			}
		}
	}()
	time.Sleep(time.Second)
	c <- 1
}

func A() int {
	fmt.Println("Start A");
	time.Sleep(time.Second*1)
	fmt.Println("End A");
	return 1
}