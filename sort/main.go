package main

import "fmt"

func bubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right
	for i < j {
		if j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}
func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func erfen(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		low := 0
		high := i - 1
		mid := -1 //中间位置
		for low <= high {
			mid = (low + high) / 2
			if temp < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j := i - 1; j >= low; j-- {
			arr[j+1] = arr[j]
		}
		if low != i {
			arr[low] = temp
		}
	}
}

func main() {
	arr := []int{9, 4, 8, 6, 4, 9, 6, 7, 8}
	// bubbleSort(arr)
	// fmt.Println(arr)
	erfen(arr)
	fmt.Println(arr)
}
