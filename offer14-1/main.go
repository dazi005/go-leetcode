package main

import (
	"fmt"
	"math"
)

func cuttingRope(n int) int {
	dp := make([]int,60)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n;i++ {
		for j := 1; j <= i;j++{
			dp[i] = int(math.Max(float64(dp[i]), math.Max(float64(dp[i-j]*j), float64((i-j)*j))))
		}
	}
	return dp[n]
}

func main(){
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(10))
}