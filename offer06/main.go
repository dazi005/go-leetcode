package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}
func reversePrint(head *ListNode) []int {
	var arr []int
	//先反转链表
	befoerNode := new(ListNode)
	for ;head != nil; {
		//新建一个节点
		p := new(ListNode)
		//给节点赋值
		p.Val = head.Val
		//这个节点指向上一个节点
		p.Next = befoerNode
		//head向后移动一个
		head = head.Next
		//把p变成上一个节点
		befoerNode = p
	}
	for ;befoerNode.Next!= nil;{
		fmt.Println("+")
		arr = append(arr,befoerNode.Val)
		befoerNode = befoerNode.Next
	}
	return arr
}

func main(){
	var b *ListNode
	head := new(ListNode)
	for i:=0;i<6;i++{
		//新建节点
		p := new(ListNode)
		p.Val = i
		p.Next = nil
		head.Next = p
		head = p
		if i == 0{
			b = head
		}
	}
	fmt.Println(reversePrint(b))
}
