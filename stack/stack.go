package stack

import "sync"

// Node ...
type Node struct {
	prev *Node
	data interface{}
}

// Stack ...
type Stack struct {
	top  *Node
	size int
	lock *sync.RWMutex
}

// New ...
func New() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// GetValue ...
func (s *Stack) GetValue() interface{} {
	if s.size == 0 {
		return nil
	}
	return s.top.data
}

// Len ...
func (s *Stack) Len() int {
	return s.size
}

// Pop ...
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.size == 0 {
		return nil
	}
	tmp := s.top
	s.top = s.top.prev
	s.size = s.size - 1
	return tmp.data
}

// Push ...
func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	node := &Node{prev: s.top, data: v}
	s.top = node
	s.size++
}
