package main

import "fmt"

type humanInterface interface {
	eat() string
	play() string
}

type man struct {
	name string
}

type woman struct {
	name string
}

func (p man) eat() string {
	return "eat banana"
}

func (p man) play() string {
	return "play game"
}

func (w woman) eat() string {
	return w.name + "eat banana"
}

func (w woman) play() string {
	return w.name+ "play game"
}

func main() {
	//test()
	test1()
}

func humandoWhat(p humanInterface) {
	fmt.Println(p.eat());
	fmt.Println(p.play());
}

func test() {
	var human humanInterface
	human = new(man)
	fmt.Println(human.eat())
	fmt.Println(human.play())
}

func test1() {
	sli := make([]humanInterface, 2)
	sli[0] = man{}
	sli[1] = woman{"woo"}
	for _, v := range sli {
		fmt.Println(v.eat())
			fmt.Println(v.play())
	}

}