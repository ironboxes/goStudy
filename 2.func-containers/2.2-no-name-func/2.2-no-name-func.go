package main

import (
	"fmt"
	"time"
)

func main() {
	// f := func(i int) {
	// 	fmt.Println(i)
	// }

	// f(1)
	// c1, c2 := closureSample(), closureSample()
	// c1()
	// c1()
	// c1()
	// 你会发现c2又从1开始输出，因为两个函数的变量是独立使用的
	// c2()
	// c2()

	test()

}

func closureSample() func() {
	count := 0
	return func() {
		count ++
		fmt.Printf("调用次数 %v \n", count)
	}
}


func test() {
	for i := 0; i < 3; i++ {
		fmt.Printf("第一次 i 产生变化中 %v \n", i)
		go func() {
			fmt.Printf("第一次输出： %v\n", i)
		}()
	}
		time.Sleep(time.Second)
}