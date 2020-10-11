package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//合并两个二叉树
//       1					   	   2                    	   3
//	 	/   \					 /   \						 /   \
//		3	  2					1      3					4     5
//     /						\     \					   / \     \
// 	  5							 4     7				 5	  4     7
func main() {
	t1 := makeTreeNode1()
	t2 := makeTreeNode2()
	res := mergeTrees(t1, t2)
	fmt.Println(res)
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return copyTreeNode(t2)
	}
	if t2 == nil {
		return copyTreeNode(t1)
	}
	newT := TreeNode{Val: t1.Val + t2.Val}
	newT.Left = mergeTrees(t1.Left, t2.Left)
	newT.Right = mergeTrees(t1.Right, t2.Right)

	return &newT
}

func copyTreeNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	var res TreeNode
	res.Val = node.Val
	res.Left = node.Left
	res.Right = node.Right
	return &res
}

func makeTreeNode1() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 1}

	TreeNodeLeft := TreeNode{Val: 3}
	TreeNodeRight := TreeNode{Val: 2}

	TreeNodeRoot.Left = &TreeNodeLeft
	TreeNodeRoot.Right = &TreeNodeRight

	TreeNodeLeftLeft := TreeNode{Val: 5}
	TreeNodeLeft.Left = &TreeNodeLeftLeft
	return &TreeNodeRoot
}

func makeTreeNode2() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 2}

	TreeNodeLeft := TreeNode{Val: 1}
	TreeNodeRight := TreeNode{Val: 3}

	TreeNodeRoot.Left = &TreeNodeLeft
	TreeNodeRoot.Right = &TreeNodeRight

	TreeNodeLeftRigth := TreeNode{Val: 4}
	TreeNodeLeft.Right = &TreeNodeLeftRigth
	TreeNodeRigthRight := TreeNode{Val: 7}
	TreeNodeRight.Right = &TreeNodeRigthRight
	return &TreeNodeRoot
}
