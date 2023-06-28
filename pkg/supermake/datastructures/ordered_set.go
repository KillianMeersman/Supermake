package datastructures

type Iter[T any] interface {
	HasNext() bool
	Next() T
	ResetIter()
}

type OrderedSet[T comparable] struct {
	elemMap   map[T]int
	elemSlice []*T
	iter_i    int
}

func NewOrderedSet[T comparable]() OrderedSet[T] {
	return OrderedSet[T]{
		elemMap:   make(map[T]int),
		elemSlice: make([]*T, 0),
	}
}

func (s OrderedSet[T]) Add(e T) {
	if !s.Contains(e) {
		s.elemMap[e] = len(s.elemSlice)
		s.elemSlice = append(s.elemSlice, &e)
	}
}

func (s OrderedSet[T]) Remove(e T) {
	index, ok := s.elemMap[e]
	if !ok {
		return
	}

	delete(s.elemMap, e)
	s.elemSlice[index] = nil
}

func (s OrderedSet[T]) Contains(e T) bool {
	_, ok := s.elemMap[e]
	return ok
}

func (s OrderedSet[T]) ExtendSlice(o []T) {
	for _, e := range o {
		s.Add(e)
	}
}

func (s OrderedSet[T]) ExtendIter(o Iter[T]) {
	for o.HasNext() {
		s.Add(o.Next())
	}
}

func (s OrderedSet[T]) Clear() {
	for e := range s.elemMap {
		delete(s.elemMap, e)
	}
	s.elemSlice = make([]*T, 0)
}

func (s OrderedSet[T]) HasNext() bool {
	for s.iter_i < len(s.elemSlice) {
		e := s.elemSlice[s.iter_i]
		s.iter_i++
		if e != nil {
			return true
		}
	}
	return false
}

func (s OrderedSet[T]) Next() T {
	return *s.elemSlice[s.iter_i]
}

func (s OrderedSet[T]) ResetIter() {
	s.iter_i = 0
}
