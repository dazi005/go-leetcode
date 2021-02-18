package main

import "fmt"

type ListNode struct{
   Val int
   Next *ListNode
}

func hasCycle1( head *ListNode ) bool {
   if head==nil||head.Next==nil{
      return false
   }
   var slow,fast = head,head.Next
   for fast != slow{
      if fast==nil||fast.Next==nil{
         return false
      }
      slow=slow.Next
      fast=fast.Next.Next
   }
   return true
}

func hasCycle(head *ListNode) bool{
   if head == nil || head.Next == nil{
      return false
   }
   slow,fast := head,head.Next
   for fast != slow{
      if slow == nil || fast.Next == nil {
         return false
      }
      slow = slow.Next
      fast = fast.Next.Next
   }
   return true
}

func main()  {
	head := &ListNode{}
	temp := head
	for i := 0; i < 10; i ++ {
       temp.Next = &ListNode{
	      Val: i,
          Next: &ListNode{},
       }
       temp = temp.Next
    }
    // 人为构造一个环
    //temp.Next = head
    fmt.Println(hasCycle(head))
}
