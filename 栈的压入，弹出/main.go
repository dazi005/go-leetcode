package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	var t = 0
	for _, v := range pushed {
		stack = append(stack, v)
		for stack[len(stack)-1] == popped[t] {
			t++
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}

	}
	if t == len(popped) {
		return true
	}
	return false
}
func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
