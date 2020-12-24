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

func suitableIndex(list []int, start int, end int, current int) int {
	// 比到最后美的比的时候就去对比下当前位置与待排元素的大小，并返回较大值的位置
	if start >= end {
		if list[start] < list[current] {
			return current
		} else {
			return start
		}
	}
	center := (end-start)/2 + start
	// 如果中间的元素比较大，就继续向左侧寻找。反之则向右
	if list[center] > list[current] {
		return suitableIndex(list, start, center, current)
	} else {
		return suitableIndex(list, center+1, end, current)
	}

}

func ErFen(list []int){
	for i := 1; i < len(list); i++ {
		// 利用二分查找，在待排元素左侧找到合适的插入位置
		p := suitableIndex(list, 0, i-1, i)
		// 如果最合适的位置不是待排元素当前位置，那就一次把合适位置后的元素都向后移动一位
		if p != i {
			temp := list[i]
			for j := i; j > p; j-- {
				list[j], list[j-1] = list[j-1], list[j]
			}
			list[p] = temp
		}
		fmt.Println(list)
	}
}
