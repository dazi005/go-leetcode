package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _,v := range matrix{
		for _,v1 := range v{
			if v1 == target{
				return true
			}
		}
	}
	return false
}

func main(){
	tArr := [][]int{
		{1,   4,  7, 11, 15},
		{2,   5,  8, 12, 19},
		{3,   6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(tArr,99))
}
