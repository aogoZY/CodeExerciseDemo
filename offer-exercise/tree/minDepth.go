package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//notice:可理解为：若求f（n）则求f（n-1），当前节点的左、右子树的最小值+1
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dl := minDepth(root.Left)
	dr := minDepth(root.Right)
	if root.Left == nil {
		return dr + 1
	} else if root.Right == nil {
		return dl + 1
	}else {
		return min(dl,dr)+1
	}
}

func main() {
	input := &TreeNode{Val: 2}
	left := &TreeNode{Val: 9}
	input.Left = left
	right := &TreeNode{Val: 20}
	input.Right = right
	res := minDepth(input)
	fmt.Println(res)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
