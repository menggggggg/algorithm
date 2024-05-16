package tree

import (
	"container/list"

	"github.com/golang-collections/collections/stack"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// 前序-递归
func PreOrder(root *TreeNode) []int {
	datas := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中
		datas = append(datas, node.Value)
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return datas
}

// 中序-递归
func InOrder(root *TreeNode) []int {
	datas := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左
		traversal(node.Left)
		// 中
		datas = append(datas, node.Value)
		// 右
		traversal(node.Right)
	}
	traversal(root)
	return datas
}

// 后续-递归
func PostOrder(root *TreeNode) []int {
	datas := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 左
		traversal(node.Left)
		// 右
		traversal(node.Right)
		// 中
		datas = append(datas, node.Value)
	}
	traversal(root)
	return datas
}

// 前序-非递归
func PreOrderNoRecursion(root *TreeNode) []int {
	datas := make([]int, 0)
	stack := stack.New()
	if root != nil {
		stack.Push(root)
	}
	for stack.Len() != 0 {
		if node := stack.Peek(); node != nil {
			stack.Pop()
			n := node.(*TreeNode)
			// 右
			if n.Right != nil {
				stack.Push(n.Right)
			}
			// 左
			if n.Left != nil {
				stack.Push(n.Left)
			}
			// 中
			stack.Push(n)
			stack.Push(nil)
		} else {
			stack.Pop()
			n := stack.Pop()
			datas = append(datas, n.(*TreeNode).Value)
		}
	}
	return datas
}

// 中序-非递归
func InOrderNoRecursion(root *TreeNode) []int {
	datas := make([]int, 0)
	stack := stack.New()
	if root != nil {
		stack.Push(root)
	}
	for stack.Len() != 0 {
		if node := stack.Peek(); node != nil {
			stack.Pop()
			// 右
			n := node.(*TreeNode)
			if n.Right != nil {
				stack.Push(n.Right)
			}
			// 中
			stack.Push(n)
			stack.Push(nil)
			// 左
			if n.Left != nil {
				stack.Push(n.Left)
			}
		} else {
			stack.Pop()
			n := stack.Pop()
			datas = append(datas, n.(*TreeNode).Value)
		}
	}
	return datas
}

// 后续-非递归
func PostOrderNoRecursion(root *TreeNode) []int {
	datas := make([]int, 0)
	stack := stack.New()
	if root != nil {
		stack.Push(root)
	}
	for stack.Len() > 0 {
		if node := stack.Peek(); node != nil {
			stack.Pop()
			// 中
			n := node.(*TreeNode)
			stack.Push(n)
			stack.Push(nil)
			// 右
			if n.Right != nil {
				stack.Push(n.Right)
			}
			if n.Left != nil {
				stack.Push(n.Left)
			}
		} else {
			stack.Pop()
			n := stack.Pop()
			datas = append(datas, n.(*TreeNode).Value)
		}
	}
	return datas
}

// 层序遍历
func LevelOrder(root *TreeNode) [][]int {
	datas := make([][]int, 0)
	list := list.New()
	if root != nil {
		list.PushBack(root)
	}
	for list.Len() > 0 {
		levelData := make([]int, 0)
		len := list.Len()
		for i := 0; i < len; i++ {
			item := list.Front()
			node := item.Value.(*TreeNode)
			levelData = append(levelData, node.Value)
			list.Remove(item)
			if node.Left != nil {
				list.PushBack(node.Left)
			}
			if node.Right != nil {
				list.PushBack(node.Right)
			}
		}
		datas = append(datas, levelData)
	}
	return datas
}
