package main

import (
	"go-leetcode/common"
)

func removeDuplicateNodes(head *common.ListNode) *common.ListNode {
	memo := make(map[int]struct{}, 0)
	result := &common.ListNode{}
	preHead := result
	for head != nil {
		if _, ok := memo[head.Val]; !ok {
			memo[head.Val] = struct{}{}
			preHead.Next = head
			preHead = head
		} else {
			preHead.Next = head.Next
		}
		head = head.Next
	}
	return result.Next
}

func main() {

}
