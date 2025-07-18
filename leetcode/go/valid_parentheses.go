package main

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		return 0
	} else {
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return item
	}
}

func (s *Stack) Peek() rune {
	if len(s.items) == 0 {
		return -1
	} else {
		item := s.items[len(s.items)-1]
		return item
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func IsValid(s string) bool {
	data := Stack{
		items: []rune{},
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			data.Push(char)
			continue
		}
		if len(data.items) == 0 {
			return false
		}
		close := data.Peek()
		if close == '(' && char == ')' || close == '{' && char == '}' || close == '[' && char == ']' {
			data.Pop()
		} else {
			return false
		}
	}

	return data.IsEmpty()
}
