package util

import (
	"iter"
	"maps"
)

type Set[T comparable] interface {
	Has(T) bool
	Add(T)
	Remove(T)
	Elements() iter.Seq[T]
	Apply(func(T))
}

type HashSet[T comparable] struct {
	elements map[T]bool
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{elements: make(map[T]bool)}
}

func (s *HashSet[T]) Has(e T) bool {
	return s.elements[e]
}

func (s *HashSet[T]) Add(e T) {
	s.elements[e] = true
}

func (s *HashSet[T]) Remove(e T) {
	delete(s.elements, e)
}

func (s *HashSet[T]) Elements() iter.Seq[T] {
	return maps.Keys(s.elements)
}

func (s *HashSet[T]) Apply(f func(T)) {
	for e := range s.elements {
		f(e)
	}
}
