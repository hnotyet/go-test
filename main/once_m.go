package main

import (
	"fmt"
	"sync"
)


type Hello struct {
	name string
}
var once sync.Once
var hello Hello
func main() {

	once.Do(func() { fmt.Println("aaaaa") })
	fmt.Println(hello.name)
	fmt.Println(hello)
}
