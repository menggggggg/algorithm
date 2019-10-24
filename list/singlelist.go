package list

// Node 节点
type Node struct {
	data interface{}
	next *Node
}

// New ...
func New(data interface{}, next *Node) *Node {
	return &Node{data: data, next: next}
}

// Data ...
func (n *Node) Data() interface{} {
	return n.data
}

// Next ...
func (n *Node) Next() *Node {
	return n.next
}

// Reverse ...
func Reverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	left := new(Node)
	current := head
	right := head.next

	for right != nil {
		current.next = left
		temp := right.next
		right.next = current
		left = current
		current = right
		right = temp
	}
	return current
}
