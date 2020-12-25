package main

import (
	"fmt"
	"math"
)

const MAX_SIZE = 1014
const STR_SIZE = 1024
/*
	二叉树
 */
type Tree struct {
	Value uint8
	Left *Tree
	Right *Tree
}

var (
	root *Tree
	list = []byte{'H','D','A','*','*','C','*','B','*','*','G','F','*','E','*','*','*'}
	index = 0
)
//先序号遍历
func (t *Tree) PreOrder() {
	if t == nil{
		return
	}
	fmt.Print( fmt.Sprintf("%c",t.Value)," ")
	t.Left.PreOrder()
	t.Right.PreOrder()
}

//中序遍历
func (t *Tree) InOrder(){
	if t == nil{
		return
	}
	t.Left.InOrder()
	fmt.Print(fmt.Sprintf("%c",t.Value)," ")
	t.Right.InOrder()
}

//后序遍历
func (t *Tree) PostOrder(){
	if t == nil{
		return
	}
	t.Left.InOrder()
	t.Right.InOrder()
	fmt.Print(fmt.Sprintf("%c",t.Value)," ")
}

//
//构造树
func BuildTree() *Tree{
	//fmt.Println(fmt.Sprintf("%c",list[index]),"   ",index)
	if list[index] == '*'{
		index ++
		return nil
	}
	tree := new(Tree)
	tree.Value = list[index]
	index ++
	tree.Left = BuildTree()
	tree.Right = BuildTree()
	return tree
}

//层次遍历
func LevelNode(t *Tree){
	q := Quene{}
	InitQueue(&q)
	if t!=nil{
		EnQueue(&q,t)
	}
	for ;!EmptyQueue(&q); {
		DeQueue(&q,t)
		fmt.Print(fmt.Sprintf("%c",t.Value)," ")
		if t.Left != nil{
			EnQueue(&q,t.Left)
		}
		if t.Right != nil{
			EnQueue(&q,t.Right)
		}
	}
}
//求深度
func (t *Tree)Depth() int{
	if t == nil{
		return 0
	}
	depthLeft := t.Left.Depth()
	depthRight := t.Right.Depth()
	return int(1 + math.Max(float64(depthLeft), float64(depthRight)))
}
//求左子叶的和
func (t *Tree)SumOfLeftLeaves() (s uint8){
	if t == nil{
		return 0
	}
	if t.Left!= nil && t.Left.Left==nil && t.Left.Right== nil {
		s = t.Left.Value
	}
	return s + t.Left.SumOfLeftLeaves() + t.Right.SumOfLeftLeaves()
}

func main(){
	root := BuildTree()
	fmt.Println("先序遍历")
	root.PreOrder()
	fmt.Println("\n中序遍历")
	root.InOrder()
	fmt.Println("\n后序遍历")
	root.PostOrder()
	fmt.Println("\n层次遍历")
	LevelNode(root)
	//var q Quene
	//if EmptyQueue(&q){
	//	InitQueue(&q)
	//	tree := Tree{
	//		Value: 1,
	//		Left: nil,
	//		Right: nil,
	//	}
	//	tree2 := Tree{}
	//	if !EnQueue(&q,&tree){
	//		fmt.Println("入队失败")
	//	}
	//	//fmt.Println(q.Front," ",q.End," ",q.Node[0].Value)
	//	if !DeQueue(&q,&tree2){
	//		fmt.Println("出队失败")
	//	}
	//	//fmt.Println(q.Front," ",q.End," ",q.Node[0].Value)
	//	fmt.Println(tree2.Value)
	//}else{
	//	fmt.Println("不是空的")
	//}

}