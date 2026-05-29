package util

type State[T any] struct {
	value T
	observableState[T]
}

func (s *State[T]) Get() T {
	return s.value
}

func (s *State[T]) Set(data T) {
	s.value = data
	s.Notify(data)
}
