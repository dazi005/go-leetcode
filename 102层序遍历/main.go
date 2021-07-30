package main

import "go-leetcode/common"

var res [][]int

func levelOrder(root *common.TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *common.TreeNode, level int) {
	if root != nil {
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level + 1)
		dfs(root.Right, level + 1)
	}
}

func main() {

}
