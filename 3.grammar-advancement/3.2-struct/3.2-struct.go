package main

import "fmt"

type people struct {
	name string
}

func (p people) toString() {
	fmt.Println(p.name)
	fmt.Printf("p的地址 %p \n", &p)
}
func (p *people) sayHello() {
	fmt.Printf("Hello! %v \n", p.name)
	fmt.Printf("*p的地址 %p \n", p)
}
func main() {
	// p := people{"ss"}
	// p.toString()

	p1 := people{"coding3min"}
p1.sayHello()
p1.toString()
p2 := &people{"tom"}
p2.sayHello()

}