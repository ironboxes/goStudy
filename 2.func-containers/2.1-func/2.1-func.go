package main

import "fmt"

func main() {
	a, b := 1, 2

	// fmt.Printf("before a=%d  b=%d\n",a, b)
	// // test(a, b)
	// test1(&a, &b)
	// fmt.Printf("after a=%d  b=%d\n",a, b)

	funcValue(a, b, add)
	funcValue(a, b, sub)

}
func test(a, b int) {
	tem := a
	a = b
	b = tem
}

func test1(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp

}

func add(a, b int) int{
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func funcValue(a, b int , v func(i, j int) int) {
	fmt.Println(v(a,b))
}