package main

type Node struct {
	Val      int
	Children []*Node
}

//给定一个 N 叉树，返回其节点值的后序遍历。

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	for _, item := range root.Children {
		res = append(res, postorder(item)...)
	}
	res = append(res, root.Val)
	return res
}
