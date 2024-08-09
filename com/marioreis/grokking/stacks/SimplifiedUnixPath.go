package stacks

import "errors"

type Stack struct {
	items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func SimplifyPath(path string) string {
	stack := NewStack()
	lastChar := rune(path[0])

	for i, char := range path {
		if i == 0 {
			if rune(char) == '/' && lastChar != '/' {
				stack.Push(rune(char))
			}

			if rune(char) == '.' && lastChar == '.' {
				currChar, _ := stack.Top()
				for currChar == '/' || currChar == '.' {
					stack.Pop()
				}

				stack.Pop()
			}

			lastChar = rune(char)
			continue
		}

	}

	return "/"
}
