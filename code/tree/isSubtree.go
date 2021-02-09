package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}

	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)

}

func check(s *TreeNode, t *TreeNode)bool {
	if s == nil && t == nil {
		return true
	}
	if s==nil || t==nil|| s.Val!=t.Val{
		return false
	}
	return check(s.Left,t.Left) &&check(s.Right,t.Right)
}
