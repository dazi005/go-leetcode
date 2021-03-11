package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	var node *TreeNode
	for i := 0; i < len(q); i++ {
		node = q[i]
		res = append(res, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return res
}

func main() {

}
