package main

//			   	3
//				/ \
//				9  20
//				/  \
//				15   7

//	输入二叉树 求最长深度
//	[3,9,20,null,null,15,7]

//递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
