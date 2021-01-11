package main

import "fmt"

func exchange(nums []int) []int {
	a := make([]int,0)
	b := make([]int,0)
	for i := 0 ;i < len(nums); i++{
		if nums[i]%2 == 0{
			b = append(b,nums[i])
		}else{
			a = append(a,nums[i])
		}
	}
	return append(a,b...)
}

func main(){
	fmt.Println(exchange([]int{1,2,3,4,5}))
}
