package main

import "fmt"

//给定两个有序链表的头指针head1和head2，打印两个链表的公共部分。
//notice:从两个链表的头开始进行如下判断：
//
//· 如果head1的值小于head2，则head1往下移动。
//
//· 如果head2的值小于head1，则head2往下移动。
//
//· 如果head1的值与head2的值相等，则打印这个值，然后head1与head2都往下移动。
//
//· head1或head2有任何一个移动到null，则整个过程停止。

func main() {
	node1 := makeNode1()
	node2 := makeNode2()
	PrintCommonPartNode(node1, node2)

}

type Node struct {
	Val  int
	Next *Node
}

func PrintCommonPartNode(node1, node2 *Node) {
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			node1 = node1.Next
		} else if node1.Val > node2.Val {
			node2 = node2.Next
		} else {
			fmt.Println(node1.Val)
			node1 = node1.Next
			node2 = node2.Next
		}
	}

}

func makeNode1() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8

	return node1
}

func makeNode2() *Node {

	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	node9 := &Node{Val: 9}

	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9
	return node5
}
