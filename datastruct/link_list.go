package main

type Item interface {
}

type Stack struct {
	items []Item
}

// stack method
func (s *Stack) New() *Stack {
	s.items = []Item{}
	return s
}
