package datastructures

type UnorderedSet[T comparable] map[T]struct{}

func NewUnorderedSet[T comparable]() UnorderedSet[T] {
	return make(UnorderedSet[T])
}

func (s UnorderedSet[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s UnorderedSet[T]) Remove(e T) {
	delete(s, e)
}

func (s UnorderedSet[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s UnorderedSet[T]) ExtendSlice(o []T) {
	for _, e := range o {
		s.Add(e)
	}
}

func (s UnorderedSet[T]) ExtendUnorderedSet(o UnorderedSet[T]) {
	for e := range o {
		s.Add(e)
	}
}

func (s UnorderedSet[T]) ExtendIter(o Iter[T]) {
	for o.HasNext() {
		s.Add(o.Next())
	}
}

func (s UnorderedSet[T]) Clear() {
	for e := range s {
		delete(s, e)
	}
}
