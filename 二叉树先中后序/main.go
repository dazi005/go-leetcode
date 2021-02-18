package main

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

var (
	pre		=	make([]int,0)
	in		=	make([]int,0)
	last	=	make([]int,0)
)

func PreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	PreOrder(root.Left, res)
	PreOrder(root.Right, res)
}

func InOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	InOrder(root.Left, res)
	*res = append(*res, root.Val)
	InOrder(root.Right, res)
}

func LastOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	LastOrder(root.Left, res)
	LastOrder(root.Right, res)
	*res = append(*res, root.Val)
}

func threeOrders( root *TreeNode ) [][]int {
	m := make([][]int,0)
	pre := make([]int, 0)
	PreOrder(root, &pre)
	m = append(m, pre)
	in := make([]int, 0)
	InOrder(root, &in)
	m = append(m, in)
	last := make([]int, 0)
	LastOrder(root, &last)
	m = append(m, last)
	return m
}

func main(){

}