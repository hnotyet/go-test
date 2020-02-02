package test

import "testing"

type Cat struct {
	Name string
	Color string
}


type BlackCat struct {
	 Cat
	 Age int
}

func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func TestCat(t *testing.T)  {

	kitty :=NewCat("kitty")
	kitty.Color = "Black"


}

