package stack

import "fmt"

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() int {
	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return d
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Print() {
	fmt.Println(s.data)
}
