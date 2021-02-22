package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func mirror(root *TreeNode){
	if root == nil {
		return
	}
	temp := &TreeNode{}
	temp = root.Left
	root.Left = root.Right
	root.Right = temp
	mirror(root.Left)
	mirror(root.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
	mirror(root)
	return root
}

func main() {

}
