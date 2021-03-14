package main

import "fmt"

//给定一个链表的头节点head，请判断该链表是否为回文结构。
//notice：用切片接受链表的值，对切片使用前后两个索引向中间夹，有不相等的值直接返回false
//
//1->2->1，返回true。
//
//1->2->2->1，返回true。
//
//1->2->3，返回false。

func makeNode1() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 2}
	node4 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	return node1
}

type Node struct {
	Val  int
	Next *Node
}

func main() {
	node := makeNode1()
	res := IsPalindrome(node)
	fmt.Println(res)
}

func IsPalindrome(node *Node) bool {
	var nodeList []int
	for node != nil {
		nodeList = append(nodeList, node.Val)
		node = node.Next
	}
	fmt.Println(nodeList)
	i := 0
	j := len(nodeList) - 1
	for i < j {
		if nodeList[i] != nodeList[j] {
			return false
		}
		i++
		j--
	}
	return true
}
