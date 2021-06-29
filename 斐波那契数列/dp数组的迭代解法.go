package main

func fib3(n int) int{
	dp := make(map[int]int, n + 1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i ++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]%1000000007
}
