package main

import (
	"fmt"
)

func main() {
	var (
		T    int
		k    int
		n    int
		arr  []int
		temp int
	)

	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&n, &k)
		arr = make([]int, n+1)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}
		//sort.Ints(arr)
		quicksort(arr, 0, len(arr)-1)
		temp = arr[k] + 1
		if k+1 > n && temp <= n && temp >= 1 {
			fmt.Println("YES")
			fmt.Println(temp)
			continue
		}
		// 如果后面有重复的就不算
		// 如果超出了范围也不算
		if k+1 > n || arr[k] == arr[k+1] {
			fmt.Println("NO")
			continue
		}
		if temp <= n && temp >= 1 {
			fmt.Println("YES")
			fmt.Println(temp)
		} else {
			fmt.Println("NO")
		}
	}
}

func quicksort(arr []int, low, high int) {
	if low < high {
		i, j := low, high
		key := arr[low]
		for i < j {
			for i < j && arr[j] >= key {
				j--
			}
			arr[i] = arr[j]
			for i < j && arr[i] <= key {
				i++
			}
			arr[j] = arr[i]
		}
		arr[i] = key
		quicksort(arr, low, i-1)
		quicksort(arr, i+1, high)
	}
}
