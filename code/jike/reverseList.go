package main

import "fmt"

//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL

func main() {
	Five := ListNode{Val: 5}
	Four := ListNode{Val: 4, Next: &Five}
	Three := ListNode{Val: 3, Next: &Four}
	Two := ListNode{Val: 2, Next: &Three}
	One := ListNode{Val: 1, Next: &Two}
	res := reverseList(&One)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
