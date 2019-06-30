package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack(10)
	stack.Push(5)
	stack.Push(5)
	stack.Push(5)
	stack.Push(5)

	fmt.Println(stack)
}
