package test
import (
	"fmt"
	"testing"
)


func TestSlice(t *testing.T) {
	var a = make([]int, 3, 5)
	a[0] = 1;
	a[1] = 2;
	a[2] = 3;
	//a[3] = 4;
	//a[4] = 5;
	fmt.Print(a)
}