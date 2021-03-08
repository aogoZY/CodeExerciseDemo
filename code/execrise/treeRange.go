package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//二叉树的前序、中序、后序遍历
func main() {
	tree := makeNode()
	//var res []int
	//preOrder(tree, &res)
	//midOrder(tree, &res)
	//afterOrder(tree, &res)
	BFS(tree)
	//bfs(tree)

}

func makeNode() *Node {
	seed6 := Node{Val: 11}
	seed5 := Node{Val: 7}
	seed4 := Node{Val: 3}
	seed3 := Node{Val: 1}
	seed2 := Node{Val: 9, Left: &seed5, Right: &seed6}
	seed1 := Node{Val: 2, Left: &seed3, Right: &seed4}
	root := Node{Val: 5, Left: &seed1, Right: &seed2}
	return &root
}

func preOrder(tree *Node, res *[]int) {
	if tree != nil {
		*res = append(*res, tree.Val)
		preOrder(tree.Left, res)
		preOrder(tree.Right, res)
	}

}

func midOrder(tree *Node, res *[]int) {
	if tree != nil {
		midOrder(tree.Left, res)
		*res = append(*res, tree.Val)
		midOrder(tree.Right, res)
	}
}

func afterOrder(tree *Node, res *[]int) {
	if tree != nil {
		afterOrder(tree.Left, res)
		afterOrder(tree.Right, res)
		*res = append(*res, tree.Val)

	}
}

//广度优先搜索 层次遍历 向波纹一样一层层的遍历
func bfs(root *Node) {
	if root == nil {
		return
	}
	// for root 需要借助队列
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:] // 通过这样的方式达到出队列
	}
}

func BFS(root *Node) {
	if root == nil {
		return
	}
	//需要借助队列
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		//var thisLineRes []int
		//thisLineRes = append(thisLineRes, node.Val)
		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		//pop出队列
		queue = queue[1:]
	}
}
