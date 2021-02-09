package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func preOrder(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		preOrder(root.Left, res)
		preOrder(root.Right, res)
	}
}

func inorder(root *TreeNode, res *[]int) {
	if root != nil {
		inorder(root.Left, res)
		*res = append(*res, root.Val)
		inorder(root.Right, res)
	}
}

func postorder(root *TreeNode, res *[]int) {
	if root != nil {
		postorder(root.Left, res)
		postorder(root.Right, res)
		*res = append(*res, root.Val)
	}
}

func MakeTreeNode() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 5}

	TreeNodeLeft := TreeNode{Val: 1}
	TreeNodeRight := TreeNode{Val: 4}

	TreeNodeRoot.Left = &TreeNodeLeft
	TreeNodeRoot.Right = &TreeNodeRight
	TreeNodeRigthLeft := TreeNode{Val: 3}
	TreeNodeRight.Left = &TreeNodeRigthLeft

	TreeNodeRigthRight := TreeNode{Val: 6}
	TreeNodeRight.Right = &TreeNodeRigthRight
	return &TreeNodeRoot
}

func main() {
	input := MakeTreeNode()
	var res []int
	postorder(input,&res)
	fmt.Println(res)
}
