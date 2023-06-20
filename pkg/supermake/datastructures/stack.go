package datastructures

type Stack[T any] struct {
	stack []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		stack: make([]T, 0),
	}
}

func (s *Stack[T]) Push(e T) {
	s.stack = append(s.stack, e)
}

func (s *Stack[T]) Pop(defaultValue T) T {
	e := s.Peek(defaultValue)
	if len(s.stack) > 0 {
		s.stack = s.stack[0 : len(s.stack)-1]
	}
	return e
}

func (s *Stack[T]) Peek(defaultValue T) T {
	if len(s.stack) < 1 {
		return defaultValue
	}
	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) Count() int {
	return len(s.stack)
}
