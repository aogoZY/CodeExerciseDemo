package main

import (
	"fmt"
	_ "fmt"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//		3
//		/ \
//		9  20
//			/  \
//			15   7

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	res := sum(root, &ans)
	fmt.Println(res)
	return ans
}

func sum(root *TreeNode, count *int) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*count += root.Left.Val
	}
	sum(root.Left, count)
	sum(root.Right, count)
	return *count
}

func main() {
	tree := makeTreeNode1()
	sumOfLeftLeaves(tree)
}

func makeTreeNode1() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 3}

	TreeNodeLeft := TreeNode{Val: 9}
	TreeNodeRight := TreeNode{Val: 20}

	TreeNodeRoot.Left = &TreeNodeLeft
	TreeNodeRoot.Right = &TreeNodeRight

	TreeNodeRightLeft := TreeNode{Val: 15}
	TreeNodeRight.Left = &TreeNodeRightLeft

	TreeNodeRightRight := TreeNode{Val: 7}
	TreeNodeRight.Right = &TreeNodeRightRight

	return &TreeNodeRoot
}
