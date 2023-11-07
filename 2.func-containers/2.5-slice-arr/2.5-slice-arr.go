package main

import "fmt"

func main() {
	//test()
	//test1()
	test2()
}

func test() {
	sli1 := []int{1, 2, 3, 4, 5}
	sli2 := make([]int, 3, 5)
	fmt.Println("sli1",sli1)
	fmt.Println(" ")
	fmt.Println("before sli12",sli2,len(sli2),cap(sli2))
	sli2 = append(sli2, 4)
	//sli2 = append(sli2, 5)
	//sli2 = append(sli2, 6)
	fmt.Println("after sli12",sli2,len(sli2),cap(sli2))
}

func test1() {
	sli := []int{0,0,0,4,5,6}
	sli1 := sli[0:1]
	fmt.Println("sli",sli,len(sli),cap(sli))
	fmt.Println("sli1",sli1,len(sli1),cap(sli1))
	sli[0] = 1
	sli1[0] = 2
	fmt.Println("sli",sli,len(sli),cap(sli))
	fmt.Println("sli1",sli1,len(sli1),cap(sli1))

}

func test2() {
	slice2 := []int{0, 0, 0, 1, 2, 3}
	slice3 := make([]int, 1, 1)
	copy(slice3, slice2[0:1])
	fmt.Println("sli2",slice2,len(slice2),cap(slice2))
	fmt.Println("sli3",slice3,len(slice3),cap(slice3))
}