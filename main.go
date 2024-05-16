package main

import (
	"fmt"

	"github.com/menggggggg/algorithm/tree"
)

func main() {
	// node5 := &list.ListNode{
	// 	Val:  5,
	// 	Next: nil,
	// }
	// node4 := &list.ListNode{
	// 	Val:  4,
	// 	Next: node5,
	// }
	// node3 := &list.ListNode{
	// 	Val:  3,
	// 	Next: node4,
	// }
	// node2 := &list.ListNode{
	// 	Val:  2,
	// 	Next: node3,
	// }
	// node1 := &list.ListNode{
	// 	Val:  1,
	// 	Next: node2,
	// }

	// newHead := list.SwapPairs(node1)
	// for tmp := newHead; tmp != nil; tmp = tmp.Next {
	// 	fmt.Print(tmp.Val)
	// 	fmt.Print(" ")
	// }
	// stack := stack.NewStack()
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Pop())

	node3 := &tree.TreeNode{
		Value: 3,
	}
	node2 := &tree.TreeNode{
		Value: 2,
	}
	node := &tree.TreeNode{
		Value: 1,
		Left:  node2,
		Right: node3,
	}
	fmt.Println(tree.LevelOrder(node))
}
