package list

import (
	"fmt"
	"strings"
)

type Array struct {
	data []interface{}
	size int
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) Add(data interface{}) {
	if len(a.data) == a.size-1 {
		a.resize()
	}
	a.data[a.size] = data
	a.size++
}

func (a *Array) Insert(data interface{}, index int) {
	if index >= a.size || index < 0 {
		panic("index out of range")
	}

	if index+1 >= len(a.data) {
		a.resize()
	}
	for i := a.size - 1; i < index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = data
}

func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(data interface{}) interface{} {
	return a.Remove(a.Find(data))
}

func (a *Array) Remove(index int) interface{} {
	if index >= a.size || index < 0 {
		panic("index out of range")
	}
	data := a.data[index]
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	return data
}

func (a *Array) Get(index int) interface{} {
	if index >= a.size || index < 0 {
		panic("index out of range")
	}
	return a.data[index]
}

func (a *Array) Set(data interface{}, index int) {
	if index >= a.size || index < 0 {
		panic("index out of range")
	}
	a.data[index] = data
}

func (a *Array) Contains(data interface{}) bool {
	for _, v := range a.data {
		if v == data {
			return true
		}
	}
	return false
}

func (a *Array) Find(data interface{}) int {
	for i, v := range a.data {
		if v == data {
			return i
		}
	}
	return -1
}

func (a *Array) String() string {

	builder := strings.Builder{}
	builder.WriteString("[")

	for i := 0; i < a.size; i++ {
		builder.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != a.size-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString("]")

	return builder.String()
}

func (a *Array) resize() {
	newData := make([]interface{}, len(a.data)*2)
	copy(newData, a.data)
	a.data = newData
	newData = nil
}

func NewArray(cap int) *Array {
	if cap < 0 {
		cap = 10
	}
	return &Array{
		data: make([]interface{}, cap),
		size: 0,
	}
}
