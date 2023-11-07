package main

import (
	"math/rand"
	"fmt"
	"sync"
)

func main() {
	//pendingforever()
	//normal()
	//standard()
	//test1()
	test2()
}

func pendingforever() {
	c := make(chan int)
	c <- 1
	a := <-c
	fmt.Println(a)
}

func normal() {
	c := make(chan int)
	go func ()  {
		c<-1
	}()
	a := <-c
	fmt.Println("a===",a)
}

func standard() {
	c := make(chan int)
	go func() {
		defer close(c)
		produceData := []int{1,2,3}
		for _, v := range produceData {
		c<-v
		}
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func test1() {
	var chanInt chan int = make(chan int, 10)
    go func() {
        defer fmt.Println("chanInt is closed")
        defer close(chanInt)
        chanInt <- 1
    }()
    res := <-chanInt
    fmt.Println(res)
}


func send(c chan<-int, wg *sync.WaitGroup) {
	//rand.Seed(time.Now().UnixNano())
	c <- rand.Int()
	wg.Done()
}

func received( c <-chan int, wg *sync.WaitGroup) {
	for gotData:= range c {
		fmt.Println(gotData)
	}
	wg.Done()
}

func test2() {
	chanInt := make(chan int , 10)
	done := make(chan struct{})
	defer close(done)
	go func ()  {
		var wg sync.WaitGroup
		defer close(chanInt)
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go send(chanInt, &wg)
		}
		wg.Wait()
	}()

	go func ()  {
		var wg sync.WaitGroup
		for i:=0; i <8; i++ {
			wg.Add(1)
			go received(chanInt,&wg)
		}
		wg.Wait()
		done <- struct{}{}
	}()

	<-done
}