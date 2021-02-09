package main

//        4					   	  4
//	 	/   \					 /   \
//		2	  7					7      2
//     / \	  / \			   / \    / \
// 	  1	 3	  6	 9	          9   6   3  1

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var res TreeNode
	res.Val=root.Val
	res.Right = invertTree(root.Left)
	res.Left = invertTree(root.Right)
	return &res
}
