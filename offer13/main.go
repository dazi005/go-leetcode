package main

import "fmt"

func movingCount(m int, n int, k int) int {
	dp := make([][]int,m + 1)
	for i := range dp{
		dp[i] = make([]int,n + 1)
	}
	return dfs(0,0,k,m,n,dp)
}

func dfs(i,j,k,m,n int,dp [][]int) int{
	if i < 0 || j < 0 || i >= m || j >= n || shuweihe(i) + shuweihe(j) > k || dp[i][j] == 1 {
		return 0
	}
	dp[i][j] = 1
	sum := 1
	sum += dfs(i + 1,j,k,m,n,dp)
	sum += dfs(i,j + 1,k,m,n,dp)
	sum += dfs(i - 1,j,k,m,n,dp)
	sum += dfs(i,j - 1,k,m,n,dp)
	return sum
}

func shuweihe(n int) (sum int){
	for n != 0{
		sum += n%10
		n = n/10
	}
	return
}

func main()  {
	fmt.Println(movingCount(2,3,1))
}
