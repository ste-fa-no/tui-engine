package util

type Observable[T any] interface {
	AddObserver(Observer[T])
	RemoveObserver(Observer[T])
	Notify(T)
}

type observableState[T any] struct {
	observers Set[Observer[T]]
}

func (s *observableState[T]) AddObserver(o Observer[T]) {
	if s.observers == nil {
		s.observers = NewHashSet[Observer[T]]()
	}
	s.observers.Add(o)
}

func (s *observableState[T]) RemoveObserver(o Observer[T]) {
	if s.observers == nil {
		s.observers = NewHashSet[Observer[T]]()
		return
	}
	s.observers.Remove(o)
}

func (s *observableState[T]) Notify(data T) {
	if s.observers == nil {
		return
	}
	s.observers.Apply(func(o Observer[T]) {
		o.Update(data)
	})
}

type Observer[T any] interface {
	Update(T)
}
