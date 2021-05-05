package main

import "go-leetcode/common"

func levelOrder(root *common.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := []*common.TreeNode{root}
	for i := 0; i < len(queue); i ++ {
		node := queue[i]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func main() {

}