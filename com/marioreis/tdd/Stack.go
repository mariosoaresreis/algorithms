package tdd

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

func (s *Stack) Pop() int {
	s.size--
	e := s.elements[s.size]
	s.elements = s.elements[0:s.size]

	return e
}
