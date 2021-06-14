package main

import "go-leetcode/common"

var skip int
var res int

func kthLargest(root *common.TreeNode, k int) int {
	skip = k
	res = 0
	dfs(root)
	return res
}

func dfs(root *common.TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	skip --
	if skip == 0 {
		res = root.Val
		return
	}
	dfs(root.Left)
}

func main() {

}
