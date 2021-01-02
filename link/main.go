package main

import (
	"fmt"
)

type ListNode struct {
	Val int
    Next *ListNode
}

//删除节点
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head} // 假头
	first := dummy // 双指针1
	second := dummy.Next // 双指针2
	for {
		if second.Val == val {
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}

	return dummy.Next
}

//用快慢指针判断是否有环
func hasCycle(head *ListNode)bool{
	if head == nil || head.Next == nil{
		return false
	}
	slow,fast := head,head.Next
	for slow != fast{
		if fast == nil || fast.Next == nil{
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

//找到入环的节点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}
	slow,fast := head,head
	for true{
		if fast == nil || fast.Next == nil{
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			break
		}
	}
	fast = head
	for slow != fast{
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
//计算链表的长度
func countLink(head *ListNode) (count int){
	if head == nil || head.Next == nil{
		return 0
	}
	slow,fast := head,head
	for true{
		if fast == nil || fast.Next == nil{
			return 0
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			break
		}
	}
	fast = head
	for slow != fast{
		slow = slow.Next
		fast = fast.Next
		count ++
	}
	for true{
		slow = slow.Next
		count ++
		if slow == fast{
			return
		}
	}
	return
}
//计算环上的节点数量
func countCycle(head *ListNode) (count int){
	if head == nil || head.Next == nil{
		return 0
	}
	slow,fast := head,head
	for true{
		if fast == nil || fast.Next == nil{
			return 0
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			break
		}
	}
	fast = head
	for slow != fast{
		slow = slow.Next
		fast = fast.Next
	}
	for true{
		slow = slow.Next
		count ++
		if slow == fast{
			return
		}
	}
	return
}

//求最远节点
func furthestNode(keyNode *ListNode) *ListNode{
	fast,slow := keyNode,keyNode
	for fast != keyNode{
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

//判断是否相交
func checkCross(headA,headB *ListNode) bool{
	headAOrigin := headA
	//先把其中一个首尾相连
	for headA.Next != nil{
		headA = headA.Next
	}
	headA.Next = headAOrigin
	if hasCycle(headB){
		return true
	}
	return false
}

func main(){
	//构建一个链表
	root := new(ListNode)
	p := new(ListNode)
	t := new(ListNode)
	single := new(ListNode)
	for i:=0;i<10;i++{
		t := new(ListNode)
		t.Val = i
		t.Next = nil
		p.Next = t
		p = t
		if i == 0{
			root = p
			*single = *p
		}
	}

	*t = *root
	for i:=0;i<5;i++{
		//fmt.Println(t.Val)
		t = t.Next
	}
	p.Next = t
	fmt.Println(countLink(root))
}
