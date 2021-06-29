package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := 999999
	for _, coin := range coins {
		subProblem := coinChange(coins, amount - coin)
		if subProblem == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(1 + subProblem)))
	}
	if res != 999999 {
		return res
	} else {
		return -1
	}
}

func main(){
	fmt.Println(coinChange1([]int{1, 2, 5}, 11))
}
