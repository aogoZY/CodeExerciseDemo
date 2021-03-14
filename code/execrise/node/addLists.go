package main

import (
	"fmt"
)

//两个链表顺序相加求和，生成新的链表
//输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912

func main() {
	node1 := makeNode1()
	node2 := makeNode2()
	res := addListReverse(node1, node2)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}

func addListReverse(node1, node2 *Node) *Node {
	res := &Node{0, nil}
	//	此处用head来指向生成链表的头节点
	head := res
	var temp int
	for node1 != nil || node2 != nil || temp > 0 {
		if node1 != nil {
			temp += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			temp += node2.Val
			node2 = node2.Next
		}
		// 链表的下一个节点为此节点两数和的余树
		res.Next = &Node{temp % 10, nil}
		//res节点后移
		res = res.Next
		//进位数
		temp = temp / 10
	}

	return head.Next
}

type Node struct {
	Val  int
	Next *Node
}

func makeNode1() *Node {
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 1}
	node3 := &Node{Val: 6}
	node1.Next = node2
	node2.Next = node3
	return node1
}

func makeNode2() *Node {
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 9}
	node7 := &Node{Val: 2}

	node5.Next = node6
	node6.Next = node7
	return node5
}


func addTwoNumbers(l1 *Node, l2 *Node) *Node {
	list := &Node{0, nil}
	//这里用一个result，用于标示生成lian bi o a d
	result := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &Node{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return result.Next
}
