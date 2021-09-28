package main

import (
	"container/list"
	"examples/datastruct/stack"
	"fmt"
)

type Node struct {
	Val   interface{}
	left  *Node
	right *Node
}

// 递归方法: 先序
func TraverseFirstNode(head *Node) {
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
func TraverseMiddleNode(head *Node) {
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
func TraverseBackNode(head *Node) {
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
// 非递归先序遍历，先压右节点，再压入左节点
func firstWorkNode(head *Node) {
	// 准备一个栈
	var stack = stack.NewStack()
	if head != nil {
		stack.Push(head)
		for !stack.IsEmpty() {
			head := stack.Pop()
			if node, ok := head.(*Node); ok {
				fmt.Println(node.Val)
				if node.right != nil {
					stack.Push(node.right)
				}
				if node.left != nil {
					stack.Push(node.left)
				}
			}
		}
	}
}

// 后序遍历：准备两个栈，中右左 => 左右中
func backWorkNode(head *Node) {
	// 准备一个栈
	var stack1 = stack.NewStack()
	var stack2 = stack.NewStack()
	if head != nil {
		stack1.Push(head)
		for !stack1.IsEmpty() {
			head := stack1.Pop()
			// 类型断言
			if node, ok := head.(*Node); ok {
				// 入栈
				stack2.Push(node)
				if node.left != nil {
					stack1.Push(node.left)
				}
				if node.right != nil {
					stack1.Push(node.right)
				}
			}
		}
	}
	for !stack2.IsEmpty() {
		pop := stack2.Pop()
		if node, ok := pop.(*Node); ok {
			fmt.Println(node.Val)
		}
	}
}

// 中序遍历，非递归，先将树的边界左节点进栈
func middleWorkNode(head *Node) {
	var stack1 = stack.NewStack()
	for !stack1.IsEmpty() || head != nil {
		if head != nil {
			// 将所有左节点进栈
			stack1.Push(head)
			head = head.left
		} else {
			// 弹出栈并打印
			n := stack1.Pop()
			if node, ok := n.(*Node); ok {
				// 判断右节点是存在, 如果不存在再次进入分支并循环
				fmt.Println(node.Val)
				head = node.right
			}
		}
	}
}

// 层序遍历， 优先遍历
func widthPucr(head *Node) {
	if head == nil {
		return
	}
	// 双向链表作为队列
	var queue = list.New()
	// map保存每个节点所在的层级
	var levelMap = make(map[interface{}]int)
	// 队列
	queue.PushBack(head)
	curlevel := 1
	levelMap[head] = curlevel
	for queue.Len() != 0 {
		cur := queue.Front()
		// 将节点从头部删除
		queue.Remove(cur)
		if n, ok := cur.Value.(*Node); ok {
			var curNodeLevel = levelMap[n]
			fmt.Println(curNodeLevel, levelMap)
			if curNodeLevel != curlevel {
				curlevel++
			}
			//fmt.Println(n.Val)
			if n.left != nil {
				queue.PushBack(n.left)
				levelMap[n.left] = curNodeLevel + 1
			}
			if n.right != nil {
				queue.PushBack(n.right)
				levelMap[n.right] = curNodeLevel + 1
			}
		}
	}
	fmt.Println(levelMap)
	fmt.Println(curlevel)
}

func main() {
	node := &Node{Val: 1}
	node.left = &Node{Val: 2}
	node.right = &Node{Val: 3}
	node.left.left = &Node{Val: 4}
	node.left.right = &Node{Val: 5}
	node.right.left = &Node{Val: 6}
	node.right.right = &Node{Val: 7}
	widthPucr(node)
}
