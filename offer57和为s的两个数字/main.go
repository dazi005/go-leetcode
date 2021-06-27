package main

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums) - 1
	for i < j {
		s := nums[i] + nums[j]
		if s < target0 {
			i ++
		} else if s > target{
			j --
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}

func main() {

}
