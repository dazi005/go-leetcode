package main

func lengthOfLongestSubstring(s string) int {
	var maxLength = 0
	hash := make(map[byte]bool, 0)
	for left, right := 0, 0 ; right < len(s) ; right ++ {
		if _, ok := hash[s[right]]; ok {
			delete(hash, s[right])
			left ++
		} else {
			hash[s[right]] = true
			right ++
		}
		if right - left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}

func maxLength( s []int ) int {
	var maxLength = 0
	hash := make(map[int]bool, 0)
	for left, right := 0, 0 ; right < len(s) ; {
		if _, ok := hash[s[right]]; ok {
			delete(hash, s[left])
			left ++
		} else {
			hash[s[right]] = true
			right ++
		}
		if right - left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}

func main(){

}
