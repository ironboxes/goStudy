package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cnt int
	shuttingDown = false
	cond = sync.NewCond(&sync.Mutex{})
)

func main() {
	//unsafeMap()
	//safeMap()
	test()
}

func unsafeMap() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	wg.Add(2)

	go func ()  {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	go func ()  {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()

	wg.Wait()


}

func safeMap() {
	var wg sync.WaitGroup
	var m sync.Map
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			m.Store(i, i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			if (i % 333 == 0) {
				fmt.Println(m.Load(i))
			}
		}
	}()
	wg.Wait()
}

func Add(entry int) {
	cond.L.Lock()
	defer cond.L.Unlock()
	cnt += entry
	fmt.Println("生产咯，来消费吧")
	cond.Signal()
}

func Get()(int, bool) {
	cond.L.Lock()
	defer cond.L.Unlock()
	for cnt == 0 && !shuttingDown {
		fmt.Println("未关闭但空了，等待生产")
		cond.Wait()
	}

	if cnt == 0 {
		fmt.Println("关闭咯，也消费完咯")
		return 0, true
	}
	cnt--
	return 1, false
}

func Shutdown() {
	cond.L.Lock()
	defer cond.L.Unlock()
	shuttingDown = true
	fmt.Println("要关闭咯，大家快消费")
	cond.Broadcast()
}



func test() {
	var wg sync.WaitGroup
	wg.Add(2)
	time.Sleep(time.Second)
	go func ()  {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			go Add(1)
			if i%5 == 0 {
				time.Sleep(time.Second)
			}
		}
	}()

	go func ()  {
		defer wg.Done()
		shuttingDown = false
		for !shuttingDown {
			var cur int
			cur, shuttingDown = Get()
			fmt.Printf("当前消费 %d, 队列剩余 %d \n", cur, cnt)
		}
	}()

	time.Sleep(time.Second * 5)
	Shutdown()
	wg.Wait()

}