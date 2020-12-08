package main

import "fmt"

func findDuplicates(nums []int) []int {
	maparr := make([]int,len(nums) + 1)
	for _,v := range nums{
		maparr[v] += 1
	}
	resultarr := make([]int,0)
	for i := 1;i < len(maparr);i++{
		if(maparr[i] == 2){
			resultarr = append(resultarr,i)
		}
	}
	return resultarr
}

func main(){
	temparr := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(temparr))
}
