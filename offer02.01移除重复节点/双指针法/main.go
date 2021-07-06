package main

import "go-leetcode/common"

func removeDuplicateNodes(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	slow := head
	for slow != nil {
		fast := slow
		for fast.Next != nil {
			if slow.Val == fast.Next.Val {
				fast.Next = fast.Next.Next
			} else {
				fast = fast.Next
			}
		}
		slow = slow.Next
	}
	return head
}

func main()  {
	
}
