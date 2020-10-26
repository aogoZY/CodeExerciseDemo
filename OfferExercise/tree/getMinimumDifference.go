package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	orderList := midOrder(root)
	res := findMin(orderList)
	return res
}

func midOrder(root *TreeNode) []int {
	var orderList []int
	midFunc(root, &orderList)
	fmt.Println(orderList)
	return orderList
}

func midFunc(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	midFunc(root.Left, res)
	*res = append(*res, root.Val)
	midFunc(root.Right, res)
}

func findMin(orderList []int) int {
	if len(orderList) < 2 {
		return orderList[0]
	}
	if len(orderList) == 2 {
		return orderList[1]-orderList[0]
	}
	small := orderList[1]-orderList[0]
	for i := 0; i < len(orderList)-1; i++ {
		if orderList[i+1]-orderList[i] < small {
			small = orderList[i+1] - orderList[i]
		}
	}
	return small
}

func main() {
	tree := makeTreeNode()
	res := getMinimumDifference(tree)
	//res:=midOrder(tree)
	fmt.Println(res)
}

func makeTreeNode() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 5}

	TreeNodeLeft := TreeNode{Val: 4}

	TreeNodeRoot.Left = &TreeNodeLeft

	TreeNodeRight := TreeNode{Val: 7}
	TreeNodeRoot.Right = &TreeNodeRight

	return &TreeNodeRoot
}
