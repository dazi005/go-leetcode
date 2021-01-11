package main

import "fmt"

func printNumbers(n int) []int {
	res := make([]int,0)
	var s = 1
	for n != 0{
		s *= 10
		n --
	}
	for i := 1; i < s ; i++{
		res = append(res,i)
	}
	return res
}

func main()  {
	fmt.Println(printNumbers(2))
}