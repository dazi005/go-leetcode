package main

import "fmt"

func findContinuousSequence(target int) [][]int{
	i, j, sum := 1, 1, 0
	res := make([][]int, 0)
	for i <= target/2 {
		if sum < target {
			sum += j
			j ++
		} else if sum > target{
			sum -= i
			i ++
		} else {
			arr := make([]int, 0)
			for k := i; k < j; k ++ {
				arr = append(arr, k)
			}
			res = append(res, arr)
			sum -= i
			i ++
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
