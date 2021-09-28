package linklist

type LinkNode struct {
	Val  interface{}
	Next *LinkNode
}

// new a test link node
func (l *LinkNode) NewLinkTestNode() *LinkNode {
	var list = []int{1, 2, 3, 2, 1}
	var pre *LinkNode = nil
	var head *LinkNode
	for i := len(list) - 1; i >= 0; i-- {
		node := &LinkNode{Val: list[i]}
		node.Next = pre
		pre = node
		if i == 0 {
			head = node
		}
	}
	return head
}
