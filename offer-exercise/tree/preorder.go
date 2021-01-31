package main

func preorder(root *Node) []int {
	if root==nil{
		return nil
	}
	var res []int
	res =append(res,root.Val)
	for _,item:=range root.Children{
		res =append(res,preorder(item)...)

	}
	return res
}
