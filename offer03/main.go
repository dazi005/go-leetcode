package main

import "fmt"

func findRepeatNumber(nums []int) int {
	var result [100000]int
	for _,v := range nums{
		result[v] ++
		if result[v] >= 2{
			return v
		}
	}
	return 0
}

func main() {
	test := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(test))
}
