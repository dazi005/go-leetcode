package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T){
	var b = []byte{'a','b','c'}
	//s := fmt.Sprintf("%c",b[0])
	//t.Log(b[0])
	//t.Log(s)
	t.Log(int('a'))
	t.Log(fmt.Sprintf("%T",b[0]))
}
