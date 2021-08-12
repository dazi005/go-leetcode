package main

import "go-leetcode/common"

/*
1. 左右节点都不是p和q的的公共祖先
2. 三种情况
*/

func lowestCommonAncestor(root, p, q *common.TreeNode) (ancestor *common.TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}

func main() {

}
