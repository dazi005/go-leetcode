package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{5,2,3,1}))
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums) - 1)
	return nums
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	i, j := low, high
	key := arr[low]
	for i < j {
		for i < j && arr[j] >= key {
			j --
		}

		arr[i] = arr[j]

		for i < j && arr[i] <= key {
			i ++
		}

		arr[j] = arr[i]
	}
	arr[i] = key
	quickSort(arr, low, i - 1)

	quickSort(arr, i + 1, high)
}

func quicksort(arr []int, low, high int) {
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
		quicksort(arr, low, i - 1)
		quicksort(arr, i + 1, high)
	}
}