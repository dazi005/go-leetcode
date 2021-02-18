package main

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

var res [][]int

func levelOrder(root *TreeNode) [][]int {
   if root == nil {
      return nil
   }
   res = make([][]int, 0)
   dfs(root, 0)
   return res
}

func dfs(root *TreeNode, level int) {
   if root == nil {
      return
   }
   if level == len(res) {
      res = append(res, []int{})
   }
   res[level] = append(res[level], root.Val)
   dfs(root.Left, level + 1)
   dfs(root.Right, level +1)
}

func main(){

}
