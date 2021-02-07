package main

import "fmt"

func main() {

	s := "abcde"

	Reverse([]byte(s))

}

func Reverse(s []byte) {
	l := len(s)

	if l == 1 {
		//fmt.Println(string(s))
		return
	}

	tmp = s[0 : l-2]

	tmp = append(tmp, s[l-1])

	fmt.Println(string(tmp))

}
