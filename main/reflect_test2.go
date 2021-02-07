package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
	fmt.Println(v.Elem().Addr())
}

