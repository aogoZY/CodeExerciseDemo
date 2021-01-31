package main

import "fmt"

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var mapCount map[int]int
	mapCount = make(map[int]int)
	findModeFunc(root, &mapCount)
	//fmt.Println(mapCount)
	res := SortMap(mapCount)
	return res

}

func SortMap(map1 map[int]int) []int {
	small := map1[0]
	var res []int
	for _, item := range map1 {
		if item > small {
			small = item

		}
	}
	for i, item := range map1 {
		if item == small {
			res = append(res, i)
		}
	}
	return res
}

func findModeFunc(root *TreeNode, mapCount *map[int]int) {
	if root == nil {
		return
	}
	(*mapCount)[root.Val] += 1
	findModeFunc(root.Left, mapCount)
	findModeFunc(root.Right, mapCount)
}

func main() {
	tree := makeTreeNode()
	res:=findMode(tree)
	fmt.Println(res)
}

func makeTreeNode() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 1}

	TreeNodeRight := TreeNode{Val: 2}

	TreeNodeRoot.Right = &TreeNodeRight

	TreeNodeRightRight := TreeNode{Val: 2}
	TreeNodeRight.Right = &TreeNodeRightRight

	return &TreeNodeRoot
}
