package tree

// 基本二叉树的实现

type Node struct {
	Val interface{}
	// 左节点
	left *Node
	// 右节点
	right *Node
}

// 二叉树的递归遍历
