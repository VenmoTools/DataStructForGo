package queue

import "testing"

func TestLinedQueue_EnQueue(t *testing.T) {
	queue := NewLoopQueue(5)
	for i := 0; i < 10; i++ {
		queue.EnQueue(i)
	}

	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()

	t.Log(queue)
}
