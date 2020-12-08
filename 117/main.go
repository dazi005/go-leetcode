package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// 数字字母映射
	mp := map[string]string{
		"1":"",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var ans []string
	var dfs func(int, string)
	// DFS函数定义
	dfs = func(i int, path string) {
		if i >= len(digits) {
			ans = append(ans, path)
			return
		}
		for _, c := range mp[string(digits[i])] {
			dfs(i + 1, path + string(c))
		}
	}
	// 执行回溯函数
	dfs(0, "")
	return ans
}

func main(){
	fmt.Println(letterCombinations("13"))
}
