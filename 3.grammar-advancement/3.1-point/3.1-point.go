package main

import "fmt"

func main() {
	//test()
	test1()
}

func test() {
	var a int
	fmt.Printf("a 的地址是：%p \n", &a)
}

func test1() {
	var a int
	var ptr *int  //一维
	var pptr **int // 二维
	var ppptr ***int // 三维

	ptr = &a
	pptr = &ptr
	ppptr = &pptr

	fmt.Printf("a的地址：%p \n", &a)
	fmt.Printf("ptr存的地址：%p \n", ptr)
	fmt.Printf("pptr存的地址的指向：%p \n", *pptr)
	fmt.Printf("ppptr存的地址的指向的指向：%p \n", **ppptr)
}