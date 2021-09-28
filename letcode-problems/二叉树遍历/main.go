package main

import "fmt"

type Node struct {
	Val interface{}
	left *Node
	right *Node
}

// 递归方法: 先序
func TraverseFirstNode(head *Node){
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	// 递归左子树
	TraverseFirstNode(head.left)
	// 递归右子树
	TraverseFirstNode(head.right)
}

// 中序遍历
func TraverseMiddleNode(head *Node){
	if head == nil {
		return
	}
	// 递归左子树
	TraverseMiddleNode(head.left)
	fmt.Println(head.Val)
	// 递归右子树
	TraverseMiddleNode(head.right)
}

// 后续遍历: 在第三次递归到本身时才打印节点
func TraverseBackNode(head *Node){
	if head == nil {
		return
	}
	// 递归左子树
	TraverseBackNode(head.left)
	// 递归右子树
	TraverseBackNode(head.right)
	fmt.Println(head.Val)
}

// 非递归方法， 压栈，出栈方式


func main(){
	node := &Node{Val: 1}
	node.left = &Node{Val: 2}
	node.right = &Node{Val: 3}
	node.left.left = &Node{Val: 4}
	node.left.right = &Node{Val: 5}
	node.right.left = &Node{Val: 6}
	node.right.right = &Node{Val: 7}
	TraverseBackNode(node)
}