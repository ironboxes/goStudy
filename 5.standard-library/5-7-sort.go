package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byAge []Person

func (b byAge) Len() int           { return len(b) }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byAge) Less(i, j int) bool { return b[i].age < b[j].age }

func main() {
	test()
}

func test() {
	peo := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(peo)
	sort.Sort(byAge(peo))
	fmt.Println(peo)
}
