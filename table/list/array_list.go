package list

import (
	"fmt"
	"strings"
)

type Array struct {
	data []int
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

func (a *Array) Add(data int) {
	if a.size >= len(a.data) {
		a.resize()
	}
	a.data[a.size] = data
	a.size++
}

func (a *Array) Insert(data int, index int) {
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

func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(data int) int {
	return a.Remove(a.Find(data))
}

func (a *Array) Remove(index int) int {
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

func (a *Array) Get(index int) int {
	if index > a.size || index < 0 {
		panic("index out of range")
	}
	return a.data[index]
}

func (a *Array) Set(data int, index int) {
	if index >= a.size || index < 0 {
		panic("index out of range")
	}
	a.data[index] = data
}

func (a *Array) Contains(data int) bool {
	for _, v := range a.data {
		if v == data {
			return true
		}
	}
	return false
}

func (a *Array) Find(data int) int {
	for i, v := range a.data {
		if v == data {
			return i
		}
	}
	return -1
}
func (a *Array) BinaryFind(data int) int {
	if !a.Sorted() {
		panic("array is not sort yet")
	}
	low := 0
	high := a.size
	for low < high {
		mid := int((low + high) / 2)
		if data > a.data[mid] {
			low = mid + 1
		} else if data < a.data[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
func (a *Array) BinaryFind2(data int) int {
	if !a.Sorted() {
		panic("array is not sort yet")
	}
	low := 0
	high := a.size
	for low < high {
		mid := int((data - a.data[low]) * (high - low) / (a.data[high] - a.data[low]))
		if data > a.data[mid] {
			low = mid + 1
		} else if data < a.data[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func (a *Array) Swap(i, j int) {
	if i < 0 || i >= a.size || j < 0 || j >= a.size {
		panic("index is illegal")
	}

	temp := a.data[i]
	a.data[i] = a.data[j]
	a.data[j] = temp

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

func (a *Array) SelectSort() {
	for i := 0; i < len(a.data); i++ {
		for j := i; j < len(a.data); j++ {
			if a.data[i] > a.data[j] {
				temp := a.data[i]
				a.data[i] = a.data[j]
				a.data[j] = temp
			}
		}
	}
}

func (a *Array) Sorted() bool {
	for i := 0; i < len(a.data)-1; i++ {
		if a.data[i] > a.data[i+1] {
			return false
		}
	}
	return true
}

func (a *Array) resize() {
	newData := make([]int, len(a.data)*2)
	copy(newData, a.data)
	a.data = newData
	newData = nil
}

func NewArray(cap int) *Array {
	if cap < 0 {
		cap = 10
	}
	return &Array{
		data: make([]int, cap),
		size: 0,
	}
}

func NewArray2(data []int) *Array {

	arr := make([]int, len(data)+5)

	copy(arr, data)

	return &Array{
		data: arr,
		size: len(arr),
	}
}
