package main

import "fmt"

/*

 */
//func BuildTreeByFront(arr []int) *Tree{
//
//}

func PreOrderTraversal(tree *Tree){
	if tree.Left != nil {
		PreOrderTraversal(tree.Left)
	}else if tree.Right != nil{
		PreOrderTraversal(tree.Right)
	}
	if tree.Left == nil || tree.Right == nil{
		fmt.Println(tree.Value)
		return
	}
}
