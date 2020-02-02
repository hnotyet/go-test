package test

import (
	"fmt"
	"testing"
)

func TestDefer(b *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
		//fmt.Printf("%d ", i)
	}
}

func TestMake(t *testing.T) {
	var a = make([]int, 3, 5)
	a[0] = 1;
	a[1] = 2;
	a[2] = 3;
	fmt.Print(a)
}
