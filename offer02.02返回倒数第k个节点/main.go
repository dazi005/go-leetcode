package main

import "go-leetcode/common"

func kthToLast(head *common.ListNode, k int) int {
	head, tail := head, head
	for i := 0; i < k; i ++ {
		tail = tail.Next
		if tail == nil {
			return 0
		}
	}
	for tail.Next != nil {
		tail = tail.Next
		head = head.Next
	}
	return tail.Val
}

func main(){

}

