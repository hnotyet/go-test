package test

import (
"regexp"
"testing"
)

func TestReg(t *testing.T) {
	reg, err := regexp.Compile("^\\d{5}$")
	if err == nil {
		flag := reg.MatchString("11111")
		print(flag)
	}
}
