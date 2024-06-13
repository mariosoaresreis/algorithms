package tdd

import "errors"

// Stack Last in first out. TDD
type Stack struct {
	elements []int
	size     int
}

func (s *Stack) Push(e int) bool {
	s.elements = append(s.elements, e)
	s.size++

	return true
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return 0, errors.New("can't pop from an empty stack")
	}

	s.size--
	e := s.elements[s.size]
	s.elements = s.elements[0:s.size]

	return e, nil
}
