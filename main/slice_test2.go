package main

import "fmt"

//func main() {
//	s := make([]int,10)
//	s = append(s, 11)
//	s = append(s, 12)
//	fmt.Println(cap(s))
//	fmt.Println(s)
//	fmt.Println(s[0])
//}

func main() {
	//a := make([]int, 10, 10)
	//FuncA(a)
	//fmt.Println(a)

	var a []int

	a[0] = 1

	fmt.Println(a)

}

func FuncA(b []int) {
	for _, v :=range b {
		v += 1
	}

}
