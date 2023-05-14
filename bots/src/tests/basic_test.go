package tests

import (
	"testing"
)

func TestBasic(t *testing.T) {
	a := 1
	b := 2
	if a+b != 3 {
		t.Error("a + b != 3")
	} else {
		t.Log("a + b == 3")
	}
}
