package queue

import "datastruct/heap"

type PriorityQueue struct {
	h *heap.MaxHeap
}

func (q *PriorityQueue) EnQueue(data int) {
	q.h.Add(data)
}

func (q *PriorityQueue) DeQueue() int {
	return q.h.ExtractMax()
}

func (q *PriorityQueue) Size() int {
	return q.h.Size()
}

func (q *PriorityQueue) Peek() int {
	return q.h.FindMax()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.h.IsEmpty()
}
