package supermake

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(e T) {
	s.stack = append(s.stack, e)
}

func (s *Stack[T]) Pop() T {
	e := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return e
}

func (s *Stack[T]) Count() int {
	return len(s.stack)
}
