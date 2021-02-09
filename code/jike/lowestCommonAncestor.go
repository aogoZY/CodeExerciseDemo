package main

import (
	"fmt"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。
//
//例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
//
//3
/// \
//5   1
/// \ / \
//6  2 0  8
/// \
//7   4

//思路：方法一：递归
// 1、判断当前节点是否和p、q相等，若相等则公共祖先为本身
//2、判断当前节点的左右子树是否可以找到p、q 若p、q分别在其左右子树在，则公共祖先为当前节点
//3、若左子树尚未找到p、q 则去右子树查找；同理，右子树找不到p、q则去左子树查找
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil && right != nil {
		return right
	}
	if left != nil && right == nil {
		return left
	}
	return nil
}

//进阶：若是二叉搜锁树则比较当前root的value和p、q的大小，若当前值大于p、q 则证明p、q在其左子树，去左子树查找
//若当前节点大于p、q则证明p、q在其右子树
//若p、q分别在当前子树左、右两边在，则证明root为其最近祖母 返回
func lowestCommonAncestorSearch(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorSearch(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorSearch(root.Right, p, q)
	}
	return root
}

func makeTreeNode() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 3}

	TreeNodeLeft := TreeNode{Val: 5}
	TreeNodeRight := TreeNode{Val: 1}

	TreeNodeRoot.Left = &TreeNodeLeft
	TreeNodeRoot.Right = &TreeNodeRight
	Tree1 := TreeNode{Val: 6}
	Tree2 := TreeNode{Val: 2}

	Tree3 := TreeNode{Val: 0}

	Tree4 := TreeNode{Val: 8}
	Tree5 := TreeNode{Val: 7}

	Tree6 := TreeNode{Val: 4}

	TreeNodeLeft.Left = &Tree1
	TreeNodeLeft.Right = &Tree2

	TreeNodeRight.Left = &Tree3
	TreeNodeRight.Right = &Tree4
	Tree1.Left = &Tree5
	Tree1.Right = &Tree6
	return &TreeNodeRoot
}
func main() {
	node := makeTreeNode()
	res := lowestCommonAncestor(node, &TreeNode{Val: 5}, &TreeNode{Val: 4})
	fmt.Println(res)
}
