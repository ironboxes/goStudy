package main

import (
	"fmt"
	"sync"
)

func main() {
	//test()
	test1()
}

func test() {
		m := make(map[string]string)
	m["name"] = "name"
	for key := range m {
		// 原来不用Printf也可以完成拼接输出啊！
		fmt.Println("key:", key, ",value:", m[key])
	}
	if v, ok := m["name"]; ok {
		fmt.Println("have ", v)
	} else {
		fmt.Println("have not ", v)
	}
	delete(m, "name")
	if v, ok := m["name"]; ok {
		fmt.Println("have ", v)
	} else {
		fmt.Println("have not ", v)
	}
}

func test1() {
	var smap sync.Map
	smap.Store("name","name")
	v, ok := smap.Load("name")
	if ok {
		fmt.Println(v)
	}
	v, ok = smap.Load("age")
	if ok {
		fmt.Println(v)
	}
		smap.Delete("age")

	smap.Range(func(key, value interface{}) bool {
		fmt.Println("key:",key,",value:",value)
		return true
	})
}