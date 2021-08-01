package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		s string
		arr []int
		k int
		lastAvg float64
		sum int
		max float64
		preIndex int
		tailIndex int
	)
	fmt.Scan(&s)
	splitDouhao := strings.Split(s, ",")
	for i := 0; i < len(splitDouhao) - 1; i ++ {
		num, _ := strconv.Atoi(splitDouhao[i])
		arr = append(arr, num)
	}
	splitMaohao := strings.Split(splitDouhao[len(splitDouhao) - 1], ":")
	num, _ := strconv.Atoi(splitMaohao[0])
	arr = append(arr, num)
	k, _ = strconv.Atoi(splitMaohao[1])
	//fmt.Println(arr, k)
	// 求上一次的平均值
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	lastAvg = float64(sum) / float64(k)

	preIndex, tailIndex = 0, k -1

	for tailIndex < len(arr) {
		sum = 0
		for i := preIndex; i <= tailIndex; i ++ {
			//fmt.Print(arr[i], " ")
			sum += arr[i]
		}
		//fmt.Println()
		avg := float64(sum) / float64(k)
		//fmt.Println(avg)
		increase := (avg -lastAvg) / lastAvg
		lastAvg = avg
		//fmt.Println(increase)
		//fmt.Println("--------")
		if increase > max {
			max = increase
		}
		preIndex ++
		tailIndex ++
	}

	fmt.Printf("%.2f", max * 100)
	fmt.Println("%")
}
