package main

import "go-leetcode/common"

func reverseList(head *common.ListNode) *common.ListNode {
	pre := &common.ListNode{}
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

func main() {

}