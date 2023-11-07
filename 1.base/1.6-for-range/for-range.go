package main

import "fmt"

func main() {
	//test2()
	//test3()
	test4()
}

func test() {
	nums := []int{1, 2, 3}
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
	}
}

func test1() {
	a, b := 1, 5
	for a < b {
		fmt.Println(a)
		a++
	}
}

func test2() {
	kvs := map[string]string{
		"a":"a",
		"b":"b",
		"c":"c",
		"d":"d",
	}
	for k, v := range kvs {
		fmt.Printf("key: %v , value: %v  \n", k, v)
	}
}

func test3() {
	tmp := []struct{
		int
		string
	}{
		{1, "a"},
		{2, "b"},
	}
	for _, v := range tmp {
		v.int = 2
	}
	for k,v := range tmp{
		fmt.Printf("k:%v, v:%v  \n",k,v)
	}
}
func test4() {
	var tmpKey *int
	var tmpValue *struct{
    int
    string
	}
	tmp := []struct{
		int
		string
	}{
		{1, "a"},
		{2, "b"},
	}
	for k,v := range tmp {
    if k == 0 {
        tmpKey = &k
        tmpValue = &v
    }
}
fmt.Println(*tmpKey)
fmt.Println(*tmpValue)
}