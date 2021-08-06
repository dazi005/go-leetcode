package main

import (
	"go-leetcode/common"
	"math"
)

// 自下而上
func isBalanced(node *common.TreeNode) bool {
	return find(node) != -1
}

func find(node *common.TreeNode) float64 {
	if node == nil {
		return 0
	}

	// 左节点
	l := find(node.Left)
	if l == -1 {
		return -1
	}

	// 右节点
	r := find(node.Right)
	if r == -1 {
		return -1
	}

	// 判断一下如果左右的差值,如果大于1，就是及时退出
	if math.Abs(l-r) > 1 {
		return -1
	}

	return math.Max(l, r) + 1
}

func main() {

}
