package main

//Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
 }
func reverseList(head *ListNode) *ListNode {
	tempHead := &ListNode{}
	beforeNode := new(ListNode)
	tempArr := []int{}
	for ;head.Next != nil;{
		tempArr = append(tempArr,head.Val)
	}
	for i := len(tempArr);i >= 0;i++ {
		nowNode := new(ListNode)
		nowNode.Val = tempArr[i]
		nowNode.Next = nil
		if i == len(tempArr){
			tempHead = nowNode
			beforeNode = nowNode
		}else {
			beforeNode.Next = nowNode
			beforeNode = nowNode
		}
	}
	return tempHead
}
