package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func makeNode1() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	return node1
}

func main() {
	node := makeNode1()
	res := reverseNode(node)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func reverseNode(node *Node) *Node {
	cur := node
	var pre *Node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		//pre, cur, cur.Next = cur, next, pre
	}
	return pre
}
