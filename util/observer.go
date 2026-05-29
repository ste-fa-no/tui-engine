package util

type Observable[T any] interface {
	Subscribe(Observer[T])
	Unsubscribe(Observer[T])
	Notify(T)
}

type ObservableState[T any] struct {
	observers Set[Observer[T]]
}

func NewObservableState[T any]() *ObservableState[T] {
	return &ObservableState[T]{observers: NewHashSet[Observer[T]]()}
}

func (s *ObservableState[T]) Subscribe(o Observer[T]) {
	s.observers.Add(o)
}

func (s *ObservableState[T]) Unsubscribe(o Observer[T]) {
	s.observers.Remove(o)
}

func (s *ObservableState[T]) Notify(data T) {
	s.observers.Apply(func(o Observer[T]) {
		o.Update(data)
	})
}

type Observer[T any] interface {
	Update(T)
}
