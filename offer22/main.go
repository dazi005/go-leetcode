package main

type ListNode struct {
    Val int
    Next *ListNode
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var start,end *ListNode
	start = head
	end = head
	for ;k>1;k--{
		end = end.Next
	}
	for end.Next != nil{
		start = start.Next
		end = end.Next
	}
	return start
}
func main(){

}
