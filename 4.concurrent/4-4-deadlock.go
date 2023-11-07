package main

import "fmt"

func main() {
	//test()
	//test1()
	//test2()
	test3()
}

func test() {
	c := make(chan int)
	c <- 1
	a := <- c
	fmt.Println(a)
}

func test1() {
	c := make(chan int)
	go func ()  {
		c <- 1
	}()
	a := <-c
	fmt.Println(a)
}

func test2() {
	c := make(chan int,1)
	c <- 1
	a := <- c
	fmt.Println(a)
}

func test3() {
	chanInt := make(chan int)
	defer close(chanInt)
  go func() {
		res := <-chanInt
		fmt.Println(res)
	}()
	chanInt <- 1
	chanInt <- 1
}

func test4() {
	chanInt := make(chan int)
	defer close(chanInt)
	go func() {
		for {
			//不使用ok会goroutine泄漏
			//res := <-chanInt
			res,ok := <-chanInt
			if !ok {
                 break
            }
			fmt.Println(res)
		}
	}()
	chanInt <- 1
	chanInt <- 1
}