package main

import "testing"

func MinString(s string, t string) string {
	if s == "" || t == "" || len(s) < len(t) {
		return ""
	}
	have := make([]int, 128)
	need := make([]int, 128)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	var (
		left  = 0
		right = 0
		start = 0
		min   = len(s) + 1
		count = 0
	)
	for right < len(s) {
		r := s[right]
		//如果不在need中继续找
		if need[r] == 0 {
			right++
			continue
		}
		if have[r] < need[r] {
			count++
		}
		have[r]++
		right++
		for count == len(t) {
			if right-left < min {
				min = right - left
				start = left
			}
			l := s[left]
			if need[l] == 0 {
				left++
				continue
			}
			if have[l] == need[l] {
				count--
			}
			left++
			have[l]--
		}
	}
	if min == len(s)+1 {
		return ""
	}
	return s[start : start+min]
}

func TestMinString(t *testing.T) {
	t.Log(MinString("ADOBECODEBANC", "ABC"))
}
