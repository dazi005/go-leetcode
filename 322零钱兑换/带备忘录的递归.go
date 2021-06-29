package main

import "math"

func coinChange1(coins []int, amount int) int {
	var memo = make(map[int]int)
	return dp(coins, amount, memo)
}


func dp(coins []int, amount int, memo map[int]int) int {
	if memo[amount] != 0 {
		return memo[amount]
	}
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := 999999
	for _, coin := range coins {
		subProblem := dp(coins, amount - coin, memo)
		if subProblem == - 1{
			continue
		}
		res = int(math.Min(float64(res), float64(subProblem+1)))
	}
	if res != 999999 {
		memo[amount] = res
	} else {
		memo[amount] = -1
	}
	return dp(coins, amount, memo)
}