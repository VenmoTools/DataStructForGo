package queue

import "datastruct/table/list"

type Queue interface {
	EnQueue(data int)
	DeQueue() int
	Size() int
	Peek() int
	IsEmpty() bool
}

type ArrayQueue struct {
	data *list.Array
}

func NewArrQueue(size int) Queue {
	return &ArrayQueue{
		data: list.NewArray(size),
	}
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.data.Size() == 0
}

func (a *ArrayQueue) EnQueue(data int) {
	a.data.Add(data)
}

func (a *ArrayQueue) DeQueue() int {
	data := a.data.Get(0)
	a.data.Remove(0)
	return data
}

func (a *ArrayQueue) Size() int {
	return a.data.Size()
}

func (a *ArrayQueue) Peek() int {
	return a.data.Get(0)
}
