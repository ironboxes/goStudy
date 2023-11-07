package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var xInt32 int32
func main() {
	//test()
	//test1()
	test2()
}

func test() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			xInt32++
		}()
	}
	wg.Wait()
	fmt.Println(xInt32)
}

func test1() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			atomic.AddInt32(&xInt32, 1)
		}()
	}
	wg.Wait()
	fmt.Println(xInt32)
}

func test2() {
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		atomic.StoreInt32(&xInt32, 100)
	}()
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&xInt32))
}