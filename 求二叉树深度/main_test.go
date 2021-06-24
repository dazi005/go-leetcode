package main

import "go-leetcode/common"

func maxDepth(root *common.TreeNode) int{
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := maxDepth(root.Left)

	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return 1 + leftDepth
	} else {
		return 1 + rightDepth
	}
}