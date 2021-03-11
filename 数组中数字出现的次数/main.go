package main

import "fmt"

func singleNumbers(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}
