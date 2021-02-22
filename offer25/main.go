package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	dom := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		}else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	}else {
		cur.Next = l1
	}
	return dom.Next
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
