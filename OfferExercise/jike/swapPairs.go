package main

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//输入：head = [1,2,3,4]
////输出：[2,1,4,3]

func swapPairs(head *ListNode) *ListNode {
	dummyHead:=&ListNode{0,head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		a := temp.Next
		b := a.Next
		temp.Next, a.Next, b.Next = b, b.Next, a
		temp = a
	}
	return dummyHead.Next
}