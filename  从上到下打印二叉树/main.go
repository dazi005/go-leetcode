package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var (
		index = 0
		arr		= make([][]int, 0)
	)
	dfs(index, root, arr)
	return arr
}

func dfs(index int, root *TreeNode, arr [][]int){
	if root == nil {
		return
	}
	if len(arr) - 1 < index {
		arr = append(arr, []int)
	}
	arr[index] = append(arr[index], root.Val)
	index ++
	dfs(index, root.Left, arr)
	dfs(index, root.Right, arr)
}

func main() {

}
