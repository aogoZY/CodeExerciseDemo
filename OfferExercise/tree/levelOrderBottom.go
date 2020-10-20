package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//3
/// \
//9  20
///  \
//15   7

//[
//[15,7],
//[9,20],
//[3]
//]
func levelOrderBottom(root *TreeNode) [][]int {
	//广度优先搜索(bfs)
	result := make([][]int, 0)
	level := 0
	if root == nil {
		return result
	}

	//初始化队列
	list := []*TreeNode{root}
	length := 1 //队列长度(即当前层节点数)
	for length > 0 {
		//从队列中取出当前层
		for i := 0; i < length; i++ {
			//出队
			node := list[0]
			list = list[1:]

			//值放入result
			if len(result) > level {
				result[level] = append(result[level], node.Val)
			} else {
				result = append(result, []int{node.Val})
			}

			//下一层入队
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}

		length = len(list)
		level++
	}

	//数组翻转
	resultLength := len(result)
	left := 0
	right := resultLength - 1
	for left < right {
		temp := result[left]
		result[left] = result[right]
		result[right] = temp

		left++
		right--
	}

	return result
}

func makeTreeNode() *TreeNode {
	TreeNodeRoot := TreeNode{Val: 3}

	TreeNodeLeft := TreeNode{Val: 9}
	TreeNodeRight := TreeNode{Val: 20}

	TreeNodeRoot.Left = &TreeNodeLeft
	TreeNodeRoot.Right = &TreeNodeRight
	TreeNodeRigthLeft := TreeNode{Val: 15}
	TreeNodeRight.Left = &TreeNodeRigthLeft

	TreeNodeRigthRight := TreeNode{Val: 7}
	TreeNodeRight.Right = &TreeNodeRigthRight
	return &TreeNodeRoot
}

//[3,9,20,null,null,15,7],
func main() {
	tree1 := makeTreeNode()
	res:=levelOrderBottom(tree1)
	fmt.Println(res)
}
