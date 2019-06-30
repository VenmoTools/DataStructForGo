package stack

import (
	"datastruct/table/list"
	"fmt"
	"strings"
)

type ListStack struct {
	list list.LinkedList
	size int
}

func (l *ListStack) Size() int {
	return l.size
}

func (l *ListStack) IsEmpty() bool {
	return l.size == 0
}

func (l *ListStack) Pop() interface{} {
	if l.size == 0 {
		panic("Empty Stack")
	}
	l.size--
	return l.list.Remove(0)
}

func (l *ListStack) Peek() interface{} {
	if l.size == 0 {
		panic("Empty Stack")
	}
	return l.list.Get(0)
}

func (l *ListStack) Push(data interface{}) {
	l.list.Insert(data.(int), l.list.Size())
	l.size++
}

func NewListStack() *ListStack {
	return &ListStack{}
}

func (s *ListStack) String() string {
	builder := strings.Builder{}
	builder.WriteString("Stack[")
	for i := 0; i < s.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", s.list.Get(i)))
		if i != s.Size()-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")
	return builder.String()
}
