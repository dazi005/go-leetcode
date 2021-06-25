package main

const MIN_VALUE = -9999

func maxSubArray(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	for j := 1; j < len(nums); j ++ {
		if dp[j - 1] > 0 {
			dp[j] = dp[j - 1] + nums[j]
		} else {
			dp[j] = nums[j]
		}
	}
	var max = MIN_VALUE
	for i := 0; i < len(dp); i ++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {

}
