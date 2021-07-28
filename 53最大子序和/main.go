package main

func maxSubArray(nums []int) int {
	// res 是目前的连续和
	max, res := nums[0], 0
	for i := 0; i < len(nums); i ++ {
		// 如果res 小于 0 就把它变成0
		if res < 0 {
			res = 0
		}
		res += nums[i]
		// 遇到新的更大的，就将其替换掉
		if res > max {
			max = res
			continue
		}
		// 如果当前的位置的值是大于0的
		if res < 0 && nums[i] > 0 {
			res = nums[i]
		}
	}
	return max
}

func main() {

}
