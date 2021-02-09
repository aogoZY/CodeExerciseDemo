package main

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root==nil{
		return 0
	}
	if root.Val<L{
		return rangeSumBST(root.Right,L,R)
	}
	if root.Val>R{
		return rangeSumBST(root.Left,L,R)
	}
	return rangeSumBST(root.Left,L,R)+rangeSumBST(root.Right,L,R)+root.Val
}

func main() {

}
