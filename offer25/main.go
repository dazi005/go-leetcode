package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	t := l
	for {
		fmt.Println(l1.Val,"   ",l2.Val)
		if l1.Val < l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
		}else {
			l.Val = l2.Val
			l2 = l2.Next
		}
		l.Next = new(ListNode)
		l = l.Next
		if l1 == nil {
			l = l2.Next
		}
		if l2 == nil {
			l = l1.Next
		}
	}
	return t
}

func main(){
	l1 := new(ListNode)
	l2 := new(ListNode)
	ll1 := l1
	ll2 := l2
	a := []int{1,2,4}
	b := []int{1,3,4}
	for i := 0;i < len(a) ; i ++ {
		l1.Val = a[i]
		l1.Next = new(ListNode)
		l1 = l1.Next
	}
	for i := 0;i < len(b) ; i ++ {
		l2.Val = a[i]
		l2.Next = new(ListNode)
		l2 = l2.Next
	}
	fmt.Println(mergeTwoLists(ll1,ll2))
}
