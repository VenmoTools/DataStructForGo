package queue

import (
	"fmt"
	"strings"
)

type LoopQueue struct {
	data []int
	font int
	tail int
	size int
}

func (l *LoopQueue) EnQueue(data int) {
	if (l.tail+1)%len(l.data) == l.font {
		l.resize(len(l.data)*2 + 1)
	}

	l.data[l.tail] = data
	l.tail = (l.tail + 1) % len(l.data)
	l.size++

}

func (l *LoopQueue) DeQueue() int {
	if l.IsEmpty() {
		panic("empty queue")
	}
	// 获取获取队首
	data := l.data[l.font]
	// 将队首位置元素清空
	l.data[l.font] = 0
	// 重新计算队首位置
	l.font = (l.font + 1) % len(l.data)
	l.size--
	if l.size == (len(l.data)-1)/4 && len(l.data)/2 != 0 {
		l.resize(len(l.data) / 2)
	}
	return data
}

func (l *LoopQueue) Size() int {
	return l.size
}

func (l *LoopQueue) Peek() int {
	if l.IsEmpty() {
		panic("empty queue")
	}
	return l.data[l.font]
}

func (l *LoopQueue) IsEmpty() bool {
	return l.font == l.tail
}

func (l *LoopQueue) String() string {
	b := strings.Builder{}
	b.WriteString("Font[")
	for i := l.font; i != l.tail; i = (i + 1) % len(l.data) {
		b.WriteString(fmt.Sprintf("%d", l.data[i]))
		b.WriteString(",")
	}
	b.WriteString("]")
	return b.String()
}

func NewLoopQueue(cap int) Queue {
	return &LoopQueue{
		data: make([]int, cap+1),
		font: 0,
		tail: 0,
	}
}

func (a *LoopQueue) resize(size int) {
	newData := make([]int, size)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[(i+a.font)%len(a.data)]
	}
	a.data = newData
	newData = nil
	a.font = 0
	a.tail = a.size
}
