package main

import "fmt"

func main() {
	node := makeNode1()
	res := removeLastKthNode(node, 8)
	fmt.Println(res)
}
func removeLastKthNode(node *Node, k int) *Node {
	slow, fast := node, node
	for k > 0 {
		fast = fast.Next
		k--
		if fast == nil && k > 0 {
			return nil
		}
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

type Node struct {
	Val  int
	Next *Node
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
