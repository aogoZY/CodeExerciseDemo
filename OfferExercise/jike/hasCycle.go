package main

//判断链表是否有环
//龟兔赛跑原则，跑得快的和跑得慢的能相遇就是有环的
func hasCycle(head *ListNode) bool {
	fast ,slow := head,head
	for slow!=nil && fast!=nil &&fast.Next!=nil{
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast{
			return true
		}
	}
	return false
}
