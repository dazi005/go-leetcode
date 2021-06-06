package main

import "go-leetcode/common"

func main() {

}

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val{
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


