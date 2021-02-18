package main

import "fmt"

func quicksort(arr []int, low, high, k int) {
	if low < high {
		i, j := low, high
		key := arr[low]
		for i < j {
			for i < j && arr[j] <= key {
				j --
			}
			arr[i] = arr[j]
			for i < j && arr[i] >= key {
				i ++
			}
			arr[j] = arr[i]
		}
		arr[i] = key
		if i == k - 1 {
			return
		}
		quicksort(arr, low, i - 1, k)
		quicksort(arr, i + 1, high, k)
	}
}

func findKth( a []int ,  n int ,  K int ) int {
	quicksort(a, 0, n - 1, K)
	return a[K - 1]
}

func main(){
	arr1 := []int{1332802,1177178,1514891,871248,753214,123866,1615405,328656,1540395,968891,1884022,252932,1034406,1455178,821713,486232,860175,1896237,852300,566715,1285209,1845742,883142,259266,520911,1844960,218188,1528217,332380,261485,1111670,16920,1249664,1199799,1959818,1546744,1904944,51047,1176397,190970,48715,349690,673887,1648782,1010556,1165786,937247,986578,798663}
	n, k := 49, 24
	fmt.Println(findKth(arr1, n , k))
	//var arr = []int{1,3,467,565,6,12,3}
	//quicksort(arr, 0, len(arr) - 1)
	//fmt.Println(arr)
}
