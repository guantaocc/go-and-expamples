package main

import (
	"examples/datastruct/linklist"
	"examples/datastruct/stack"
	"fmt"
)

type Node = *linklist.LinkNode

// 判断一个回文链表
func huiwenLinkNode(head Node) bool {
	// describe a stack
	s := &stack.Stack{}
	st := s.New()
	// 将所有链表节点入栈
	var cur = head
	for cur != nil {
		st.Push(cur)
		cur = cur.Next
	}

	var isHeadTailOnce bool = true
	// 比对节点
	for head != nil {
		// 依次比对节点
		node := st.Pop()
		// 类型断言
		if n, ok := node.(Node); ok {
			if head.Val != n.Val {
				isHeadTailOnce = false
				break
			}
		}
		head = head.Next
	}
	return isHeadTailOnce
}

func main() {
	testList := &linklist.LinkNode{}
	list := testList.NewLinkTestNode()
	fmt.Println(list)
	isHuiwen := huiwenLinkNode(list)
	fmt.Println(isHuiwen)
}
