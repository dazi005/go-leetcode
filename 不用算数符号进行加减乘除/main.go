package main

import "fmt"

func add(a int, b int) int {
	var carry int
	for b != 0 {
		carry = (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(9, 2))
}