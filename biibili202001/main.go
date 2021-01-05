package main

import "fmt"

var (
	result = make([]int,0)
)

func dfs(x int) int{
	if x == 1{
		return a(x)
	}
	if x == 2{
		return b(x)
	}
	return a(x - 1) + a(x - 2)
}

func a(x int) int{
	result = append(result,2)
	return 2 * x + 1
}

func b(x int) int {
	result = append(result,3)
	return 2 * x + 2
}

func main(){
	fmt.Println(dfs(10))
}
