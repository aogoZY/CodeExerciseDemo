package main

//
//输入：2->1->4->3  【2，1，4，3】 【1，2，3，4】
//输出：1->2->3->4

2->1->4->3

1->2->4->3


1->2->4->3

1->2->3->4
type Node struct {
	Val  int
	Next *Node
}

func SortNode(head *Node) {
	var sortFlag bool
	for !sortFlag {
		for head != nil {
			sortFlag = true
			cur := head
			if cur.Next.Val < cur.Val {
				temp := cur.Val
				cur.Val = cur.Next.Val
				cur.Next.Val = temp
				cur = cur.Next
				sortFlag = false
			}
			if sortFlag == true {
				break
			}
		}
	}
}
