package main

import "fmt"

func missingNumber(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right)/2
		if nums[mid] != mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	arr := []int{1,2,3,4,6,7,8,9,10,11}
	fmt.Println((0 + len(arr))>>1)
	fmt.Println((0 + len(arr))/2)

}
