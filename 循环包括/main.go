package main

import "fmt"

// src 是原字符串， target是要判断的字符串
func compar(src, target string) bool {
	// 把原来的字符串存到map里，比较的时候可以提高效率
	m1 := make(map[string]int, 0)
	for i := 0; i < len(src); i++ {
		m1[string(src[i])] = i
	}
	// 如果包含就开始判断
	if v, ok := m1[string(target[0])]; ok {
		t := 0
		// target的下标从0开始累加，srv的下标从map中保存的下标累加（到最大的时候重置为0）
		// 如果target走完了，那么就包含
		// 如果中途发现了不一样的，就退出循环
		for target[t] == src[v] {
			v++
			t++
			if t == len(target) {
				return true
			}
			if v == len(src) {
				v = 0
			}
		}
	}
	return false
}

func main() {
	fmt.Println(compar("aaababa", "aaaab"))
	fmt.Println(compar("abcd", "cd"))
	fmt.Println(compar("abcd", "db"))
	fmt.Println(compar("abcd", "ac"))
}
