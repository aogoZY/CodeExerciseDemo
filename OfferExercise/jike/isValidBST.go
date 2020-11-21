package main

import (
	"fmt"
	"math"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//判断某树是否为二叉搜索树
//1、将其做中序遍历后，看list是否为严格升序递增的，若是则为二叉搜索树
//2、递归查询左子树的最大值是否小于根节点,查询右子树的最小值是否大于根节点

func isValidBST(root *TreeNode) bool {
	var res []int
	inorder(root, &res)
	fmt.Println(res)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

//将树做中序排列
func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
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
	node := MakeTreeNode()
	isValidBST(node)

}

//设计一个递归函数 helper(root, lower, upper) 来递归判断，
// 表示考虑以 root 为根的子树，判断子树中所有节点的值是否都在 (l,r)(l,r) 的范围内（注意是开区间）。
// 如果 root 节点的值 val 不在 (l,r)(l,r) 的范围内说明不满足条件直接返回，
// 否则我们要继续递归调用检查它的左右子树是否满足，如果都满足才说明这是一棵二叉搜索树。

func isValidBST2(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)

}
