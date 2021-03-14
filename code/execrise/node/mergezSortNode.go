package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

//给定两个有序单链表的头节点 head1和 head2，请合并两个有序链表，合并后的链表依然有序，并返回合并后链表的头节点。
//
//0->2->3->7->null
//
//1->3->5->7->9->null
//
//合并后的链表为：0->1->2->3->3->5->7->7->9->null

//notice:
//1．如果两个链表中有一个为空，说明无须合并过程，返回另一个链表的头节点即可。
//2．比较head1和head2的值，小的节点是合并链表的头节点，记为head；
// 在之后的步骤里，哪个链表的头节点的值更小，另一个链表的所有节点都会依次插入到这个链表中。
//3．不妨设head节点所在的链表为链表1，另一个链表为链表2。链表1和链表2都从头部开始一起遍历，比较每次遍历到的两个节点的值，
// 记为cur1和cur2，然后根据大小关系做出不同的调整，同时用一个变量pre表示上次比较时值较小的节点。
//4．如果链表1先走完，此时cur1=null, pre为链表1的最后一个节点，那么就把pre的next指针指向链表2当前的节点（即cur2），表示把链表2没遍历到的有序部分直接拼接到最后，
//  如果链表2先走完，说明链表2的所有节点都已经插入到链表1中，调整结束。
//5．返回合并后链表的头节点head。

func main() {
	node1 := makeNode1()
	node2 := makeNode2()
	res := GetOrderNode(node1, node2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func GetOrderNode(node1, node2 *Node) *Node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	var head, cur1, cur2 *Node
	// 比较两个链表的头节点 值更小的为head头
	//cur1、cur2分别为两个链表的当前值，我们默认头节点所在的链表为cur1

	if node1.Val < node2.Val {
		head = node1
		cur1 = node1
		cur2 = node2

	} else {
		head = node2
		cur1 = node2
		cur2 = node1
	}

	//pre节点存放每次cur1、cur2比较的最小值
	var pre, next *Node
	for cur1 != nil && cur2 != nil {
		//如果头节点所在链表的值小，需要把小值的位置存放到pre中，同时cur1链表往后移一位
		if cur1.Val < cur2.Val {
			pre = cur1
			cur1 = cur1.Next
		} else {
			//如果非头节点所在链表的值小，需要将非头节点接入头节点，小值存放到pre中，非头节点往后移一位
			next = cur2.Next
			pre.Next = cur2
			cur2.Next = cur1 //这两步是将当前节点接入头节点所在链表
			pre = cur2
			cur2 = next
		}
	}
	if cur1 == nil {
		pre.Next = cur2
	} else {
		pre.Next = cur1
	}
	return head
}

func makeNode1() *Node {
	node1 := &Node{Val: 0}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 7}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	return node1
}

func makeNode2() *Node {
	node5 := &Node{Val: 1}
	node6 := &Node{Val: 3}
	node7 := &Node{Val: 5}
	node8 := &Node{Val: 7}
	node9 := &Node{Val: 9}

	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9
	return node5
}
