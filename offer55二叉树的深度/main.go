package main

import "go-leetcode/common"

func main() {

}

func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	lenLeft := maxDepth(root.Left)

	lenRight := maxDepth(root.Right)

	if lenLeft > lenRight {
		return lenLeft + 1
	} else {
		return lenRight + 1
	}
}

// 简化版
func maxDepth2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
