package main

import "go-leetcode/common"

var res [][]int

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res = make([][]int,1)
	subFunc(root, 0)
	return res
}

func subFunc(root *TreeNode, depth int) {
	if depth >= len(res) {
		res = append(res, []int{root.Val})
	} else {
		res[depth] = append(res[depth], root.Val)
	}
	if root.Left != nil {
		subFunc(root.Left, depth + 1)
	}
	if root.Right != nil {
		subFunc(root.Right, depth + 1)
	}
}

func main() {

}
