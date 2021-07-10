package main

import "go-leetcode/common"

func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	newList := &common.ListNode{}
	head := newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Val = l1.Val
			l1 = l1.Next
		} else {
			newList.Val = l2.Val
			l2 = l2.Next
		}
		newList.Next = new(common.ListNode)
		newList = newList.Next
	}
	if l1.Next == nil {
		newList.Next = l2
	}
	if l2.Next == nil {
		newList.Next = l1
	}
	return head
}

func main() {

}
