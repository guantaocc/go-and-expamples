package main

import "fmt"

type LinkNode struct {
	Val  interface{}
	Next *LinkNode
}

func addLinkNode() (head *LinkNode) {
	var pre *LinkNode
	for i := 5; i > 0; i-- {
		node := &LinkNode{Val: i, Next: pre}
		pre = node
		if node.Val == 1 {
			head = node
		}
	}
	return
}

// 反转链表
func reverseLinkList(head *LinkNode) *LinkNode {
	var pre *LinkNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {
	// add node
	head := addLinkNode()
	// reverse linkList
	cur := reverseLinkList(head)
	fmt.Printf("%v", cur)
}
