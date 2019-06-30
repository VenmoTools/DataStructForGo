package queue

import "datastruct/table/list"

type LinedQueue struct {
	data *list.LinkedList
}

func NewLinedQueue(size int) Queue {
	return &LinedQueue{
		data: list.NewLinkedList(),
	}
}

func (a *LinedQueue) IsEmpty() bool {
	return a.data.Size() == 0
}

func (a *LinedQueue) EnQueue(data int) {
	a.data.Add(data)
}

func (a *LinedQueue) DeQueue() int {
	data := a.data.Get(0)
	a.data.Remove(0)
	return data
}

func (a *LinedQueue) Size() int {
	return a.data.Size()
}

func (a *LinedQueue) Peek() int {
	return a.data.Get(0)
}
