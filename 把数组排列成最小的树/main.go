package main

import (
	"fmt"
	"sort"
	"strings"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i], nums[j])
	})
	fmt.Println(nums)
	var res strings.Builder
	for _, v := range nums {
		res.WriteString(fmt.Sprintf("%d", v))
	}
	return res.String()
}

func compareNumber(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	if str1 < str2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(minNumber([]int{90, 23, 3, 9}))
}
