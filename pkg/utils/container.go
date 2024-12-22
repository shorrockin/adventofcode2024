package utils

import "container/list"

type Container[T any] struct {
	list     *list.List
	popFront bool
}

func NewQueue[T any]() *Container[T] {
	return &Container[T]{list: list.New(), popFront: true}
}

func NewStack[T any]() *Container[T] {
	return &Container[T]{list: list.New(), popFront: false}
}

func (q *Container[T]) Push(value T) {
	q.list.PushBack(value)
}

func (q *Container[T]) PushAll(values ...T) {
	for _, value := range values {
		q.Push(value)
	}
}

func (q *Container[T]) Pop() (T, bool) {
	if q.list.Len() == 0 {
		var zero T
		return zero, false
	}
	var element *list.Element

	if q.popFront {
		element = q.list.Front()
	} else {
		element = q.list.Back()
	}

	q.list.Remove(element)
	return element.Value.(T), true
}

func (q *Container[T]) MustPop() T {
	value, ok := q.Pop()
	if !ok {
		panic("failed to dequeue, expected value to be on queue")
	}
	return value
}

func (q *Container[T]) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Container[T]) Len() int {
	return q.list.Len()
}

func (q *Container[T]) Clear() {
	q.list.Init()
}
