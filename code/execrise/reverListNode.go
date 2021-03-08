package main

import "fmt"

type TreeNode struct {
	Val  int
	Next *TreeNode
}

func main() {
	head := makeTreeNode()
	res := ReverseNode(head)
	for res.Next != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Println(res)
}

func makeTreeNode() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 5}

	TreeNodeLeft := TreeNode{Val: 1}
	TreeNodeRoot.Next = &TreeNodeLeft

	TreeNodeRigthLeft := TreeNode{Val: 3}
	TreeNodeLeft.Next = &TreeNodeRigthLeft

	TreeNodeRigthRight := TreeNode{Val: 6}
	TreeNodeRigthLeft.Next = &TreeNodeRigthRight
	return &TreeNodeRoot
}

func ReverseNode(head *TreeNode) *TreeNode {
	var pre *TreeNode
	for head != nil {
		head.Next, pre, head = pre, head, head.Next
	}
	return pre
}

func reverseList(head *TreeNode) *TreeNode {
	cur := head
	var pre *TreeNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
