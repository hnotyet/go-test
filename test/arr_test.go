package test

import (
	"testing"
)

func TestArr(t *testing.T) {
	arr := [4]int{1, 2, 3}

	t.Log(arr)
	t.Log(len(arr))
	t.Log(cap(arr))



}
