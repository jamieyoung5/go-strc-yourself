package queue

import (
	"container/list"
	"sync"
)

type Queue[T any] struct {
	list *list.List
	lock sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: list.New()}
}

func (q *Queue[T]) Enqueue(value T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.list.PushBack(value)
}

func (q *Queue[T]) Dequeue() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	element := q.list.Front()
	if element != nil {
		q.list.Remove(element)
		return element.Value.(T)
	}

	var zeroValue T
	return zeroValue
}

func (q *Queue[T]) Peek() T {
	q.lock.Lock()
	defer q.lock.Unlock()

	element := q.list.Front()
	if element != nil {
		return element.Value.(T)
	}

	var zeroValue T
	return zeroValue
}

func (q *Queue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.list.Len() == 0
}

func (q *Queue[T]) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.list.Len()
}
