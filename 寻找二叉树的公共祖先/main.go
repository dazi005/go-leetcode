package main

import "go-leetcode/common"

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	// 边界条件1 这是一个空树
	if root == nil {
		return root
	}

	// 边界条件2 p或者q中有一个是根结点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 从左边开始找
	left := lowestCommonAncestor(root.Left, p, q)

	// 从右边开始找
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

func main() {

}
