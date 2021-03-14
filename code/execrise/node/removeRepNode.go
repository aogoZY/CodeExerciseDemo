package main

import "fmt"

//给定一个无序单链表的头节点head，删除其中值重复出现的节点。
//
//例如：1->2->3->3->4->4->2->1->1->null，删除值重复的节点之后为1->2->3->4->null。
//
//方法1：如果链表长度为N，时间复杂度达到O(N)。
//		遍历链表，使用map判断该值是否已经存在，
//		若是cur.val已经存在，将当前节点删除，pre.next 指向cur.next,
//		若cur.val不存在，将当前值存进map，并更新pre的值。
//
//方法2：额外空间复杂度为O(1)。
//		类似排序算法，第一遍遍历结果为1->2->3->3->4->4->2->null，
//		第二遍为1->2->3->3->4->4->null，
//		第三遍为1->2->3->4->4->null，
//		第四遍为1->2->3->4->null

func main() {
	node := makeNode1()
	res := removeRepNode(node)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type Node struct {
	Val  int
	Next *Node
}

func removeRepNode(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	pre := node
	cur := node.Next
	nodeMap := make(map[int]bool)
	nodeMap[node.Val] = true
	for cur != nil {
		next := cur.Next
		_, ok := nodeMap[cur.Val]
		if ok {
			pre.Next = next
		} else {
			nodeMap[cur.Val] = true
			pre = cur
		}
		cur = next
	}
	return node
}

func removeRepNode2(node *Node) *Node {
	//cur 用于确定当前最外层遍历需要比较的节点，i
	cur := node
	//pre用于保存更新值最后的index位置
	//next用于循环内部的j
	var pre, next *Node
	for cur != nil {
		pre = cur
		next = cur.Next
		for next != nil {
			if cur.Val == next.Val {
				pre.Next = next.Next
			} else {
				pre = next
			}
			next = next.Next
		}
		cur = cur.Next
	}
	return node
}

func makeNode1() *Node {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 3}
	node5 := &Node{Val: 4}
	node6 := &Node{Val: 2}
	node7 := &Node{Val: 1}
	node8 := &Node{Val: 1}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	return node1
}

//func removeRepNode2(node *Node) *Node {
//	head := node
//	compare, pre, cur := node, node, node
//	for compare != nil {
//		if compare.Val == cur.Val {
//
//		}
//	}
//	return nil
//}
