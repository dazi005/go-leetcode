package main

import "strconv"

/*

func translateNum(num int) int {
	num_str := strconv.Itoa(num)
	res := 1
	q := 1
	p := 0
	for i := 1; i < len(num_str); i++ {
		p = q
		q = res
		sum := num_str[i-1:i+1]
		if sum >= "10" && sum <= "25" {
			res += p
		}
	}
	return res
}

作者：su-huan-zhen-fVk9kEa7uP
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/solution/golangjie-fa-by-su-huan-zhen-fvk9kea7up/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func translateNum(num int) int {
	num_str := strconv.Itoa(num)
	res := 1
	p := 0
	q := 1
	for i := 1; i < len(num_str); i++ {
		p = q
		q = res
		num_temp := num_str[i-1 : i+1]
		if num_temp <= "25" && num_temp >= "10" {
			res += p
		}
	}
}

func main() {

}
