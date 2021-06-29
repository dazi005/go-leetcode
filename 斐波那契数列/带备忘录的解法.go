package main

func fib2(n int) int {
	if n < 1 {
		return 0
	}
	memo := make(map[int]int, n + 1)
	return helper(memo, n)
}

func helper(memo map[int]int,n int) int{
	if n == 1 || n == 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n - 1) + helper(memo, n - 2)
	return memo[n]%1000000007
}