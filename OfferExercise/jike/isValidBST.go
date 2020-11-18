package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//判断某树是否为二叉搜索树
//将其做中序遍历后，看list是否为严格升序递增的，若是则为二叉搜索树
func isValidBST(root *TreeNode) bool {
	var res []int
	inorder(root, &res)
	fmt.Println(res)
	for i:=0;i<len(res)-1;i++{
		if res[i]>=res[i+1]{
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

func makeTreeNode() *TreeNode {
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
	node := makeTreeNode()
	isValidBST(node)

}
