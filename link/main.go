package main
type ListNode struct {
	Val int
    Next *ListNode
}

//删除节点
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head} // 假头
	first := dummy // 双指针1
	second := dummy.Next // 双指针2
	for {
		if second.Val == val {
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}

	return dummy.Next
}

//判断是否有环
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	nodesMap := make(map[*ListNode]bool, 0)
	for cur := head; cur != nil; cur = cur.Next {
		if !nodesMap[cur] {
			nodesMap[cur] = true
		} else {
			return cur
		}
	}
	return nil
}
func main(){

}
