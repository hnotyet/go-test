package test

import (
	"fmt"
	"testing"
)

type A struct {
	a int
}

type B struct {
	b int
}

type C struct {
	A
	B
}

//func Testaaa(){}

func TestInnerStruct(t *testing.T) {
	c := &C{}
	//c.A.a = 1
	//json.Marshal(Testaaa)
	fmt.Println(c)
}
