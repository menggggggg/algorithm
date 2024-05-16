package list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
// https://leetcode.cn/problems/reverse-linked-list/

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode

	for cur := head; cur != nil; {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

// 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func SwapPairs(head *ListNode) *ListNode {
	// 使用虚头节点
	if head == nil {
		return nil
	}
	dummyNode := &ListNode{
		Next: head,
	}
	newHead := dummyNode
	for {
		if dummyNode.Next == nil || dummyNode.Next.Next == nil {
			break
		}
		n1, n2 := dummyNode.Next, dummyNode.Next.Next
		dummyNode.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		// 更换虚节点
		dummyNode = n1
	}
	return newHead.Next
}
