package main

func main() {

}

func helper(nums []int, target int) int{
	i, j := 0, len(nums) - 1
	for i <= j {
		// 中间
		m := (i + j) /2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return i
}

func search(nums []int, target int) int {
	return helper(nums, target) - helper(nums, target - 1)
}
