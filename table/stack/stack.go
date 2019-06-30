package stack

import (
	"datastruct/table/list"
	"fmt"
	"strings"
)

type Stack interface {
	Size() int
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Peek() interface{}
}

func NewArrayStack(cap int) Stack {
	return &ArrayStack{
		arr: list.NewArray(cap),
	}
}

type ArrayStack struct {
	arr *list.Array
}

func (s *ArrayStack) Pop() interface{} {
	return s.arr.RemoveFirst()
}

func (s *ArrayStack) Push(data interface{}) {
	s.arr.Insert(data, 0)
}

func (s *ArrayStack) IsEmpty() bool {
	return s.arr.IsEmpty()
}

func (s *ArrayStack) Size() int {
	return s.arr.Size()
}

func (s *ArrayStack) Peek() interface{} {
	return s.arr.Get(0)
}

func (s *ArrayStack) String() string {
	builder := strings.Builder{}
	builder.WriteString("Stack[")
	for i := 0; i < s.Size(); i++ {
		builder.WriteString(fmt.Sprintf("%v", s.arr.Get(i)))
		if i != s.Size()-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")
	return builder.String()
}
