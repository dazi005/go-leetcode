package main

type ListNode struct {
	Val int
    Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	var before = head
	if before.Val == val{
		head = head.Next
	}
	var tHead = before.Next
	for tHead != nil{
		if tHead.Val == val{
			before.Next = tHead.Next
		}
		before = tHead
		tHead = tHead.Next
	}
	return head
}

func main(){

}
