package main

import "fmt"

func checkIsOk(arr []int, n int) bool{
	for _, v := range arr {
		if v > n {
			return false
		}
	}
	return true
}

func main(){
	var n, m, h int
	var arr = make([]int, 0)
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&h)
	for i := 0; i < n;i ++ {
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	for i := 0; i < n; i ++ {
		if n - i < m {
			break
		}
		if arr[i] > h {
			continue
		}
		if arr[i] <= h {
			if checkIsOk(arr[i: i + m], h) {
				fmt.Println(i + 1)
				return
			} else {
				continue
			}
		}
	}
	fmt.Println(-1)
}
