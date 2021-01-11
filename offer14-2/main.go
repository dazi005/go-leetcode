/*
	这一体只能用贪心了，因为涉及到大数运算
 */
package main

import (
	"fmt"
)

func cuttingRope(n int) int {
	if n < 4{
		return n - 1
	}
	if n == 4{
		return 4
	}
	var res = 1
	for n > 4{
		res *= 3
		res %= 1000000007
		n -= 3
	}
	return res * n%1000000007
}

func main(){
	fmt.Println(cuttingRope(10))
}