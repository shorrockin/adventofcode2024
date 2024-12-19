package utils

import "container/list"

type Queue[T any] struct {
	list     *list.List
	popFront bool
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: list.New(), popFront: true}
}

func NewHeap[T any]() *Queue[T] {
	return &Queue[T]{list: list.New(), popFront: false}
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) EnqueueAll(values ...T) {
	for _, value := range values {
		q.Enqueue(value)
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
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

func (q *Queue[T]) MustDequeue() T {
	value, ok := q.Dequeue()
	if !ok {
		panic("failed to dequeue, expected value to be on queue")
	}
	return value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Len() == 0
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Clear() {
	q.list.Init()
}
