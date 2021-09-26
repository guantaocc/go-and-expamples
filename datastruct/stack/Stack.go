package stack

// 一个基本的栈结构实现
type stackItem interface{}

type Stack struct {
	items []stackItem
}

// new stack
func (s *Stack) New() *Stack {
	s.items = make([]stackItem, 0)
	return s
}

// add stack item
// 从栈尾部增加元素, 入栈
func (s *Stack) Push(item ...interface{}) *Stack {
	for _, value := range item {
		s.items = append(s.items, value)
	}
	return s
}

// 出栈
func (s *Stack) Pop() interface{} {
	popItem := s.items[:1]
	s.items = s.items[1:]
	return popItem
}

// is empty
func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	} else {
		return false
	}
}
