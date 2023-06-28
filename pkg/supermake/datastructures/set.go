package datastructures

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) ExtendSlice(o []T) {
	for _, e := range o {
		s.Add(e)
	}
}

func (s Set[T]) ExtendSet(o Set[T]) {
	for e := range o {
		s.Add(e)
	}
}

func (s Set[T]) Clear() {
	for e := range s {
		delete(s, e)
	}
}
