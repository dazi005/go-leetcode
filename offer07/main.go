package main

type TreeNode struct {
	Val int
	Left *TreeNode
    Right *TreeNode
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	index := map[int]int{}
	for i,v := range inorder{
		index[v] = i
	}
	return myBuildTree(index,preorder,0,len(preorder) - 1,0,len(inorder) - 1)
}

func myBuildTree(index map[int]int,preorder []int,preStart,preEnd,inStart,inEnd int) *TreeNode{
	if len(preorder) == 0 || preEnd - preStart < 0 || inEnd - inStart < 0{
		return nil
	}
	m := index[preorder[preStart]]
	t := &TreeNode{
		Val: preorder[preStart],
		/*
			m - inStart是偏移量

			[preStart + 1,preStart + m - inStart]是根据m算出来的 左 边的树在preorder数组里面的的区间
			[inStart,m - 1]是根据m算出来的 左 边的树在inorder数组里面的区间
		*/
		Left: myBuildTree(index,preorder,preStart + 1,preStart + m - inStart,inStart,m - 1),
		/*
			[preStart + 1 + m - inStart,preEnd]是根据m算出来的 右 边的树在preorder数组里面的的区间
			[m + 1,inEnd]是根据m算出来的 右 边的树在inorder数组里面的区间
		 */
		Right: myBuildTree(index,preorder,preStart + 1 + m - inStart,preEnd,m + 1,inEnd),
	}
	return t
}