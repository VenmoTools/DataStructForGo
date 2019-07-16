package heap

import "datastruct/table/list"

type MaxHeap struct {
	array *list.Array
}

func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		array: list.NewArray(cap),
	}
}

func (m *MaxHeap) Size() int {
	return m.array.Size()
}

func (m *MaxHeap) IsEmpty() bool {
	return m.array.IsEmpty()
}

func (m *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("index 0 does not have parent")
	}

	return (index - 1) / 2
}

func (m *MaxHeap) leftChild(index int) int {

	return index*2 + 1
}

func (m *MaxHeap) rightChild(index int) int {

	return index*2 + 2
}

func (m *MaxHeap) Add(data int) {
	m.array.Add(data)
	m.siftUp(m.array.Size() - 1)
}

func (m *MaxHeap) siftUp(index int) {
	// 如果当前元素大于父元素
	for index > 0 && m.array.Get(m.parent(index)) < m.array.Get(index) {
		// 交换当前元素和父元素位置
		m.array.Swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *MaxHeap) FindMax() int {
	if m.array.Size() == 0 {
		panic("heap is empty")
	}
	return m.array.Get(0)
}

func (m *MaxHeap) ExtractMax() int {
	ret := m.FindMax()

	m.array.Swap(0, m.Size()-1)
	m.array.RemoveLast()

	m.siftDown(0)

	return ret
}

func (m *MaxHeap) siftDown(k int) {
	for m.leftChild(k) < m.array.Size() {
		j := m.leftChild(k)
		if j+1 < m.array.Size() && m.array.Get(j+1) < m.array.Get(j) {
			j = m.rightChild(k)
		}

		if m.array.Get(k) >= m.array.Get(j) {
			break
		}

		m.array.Swap(k, j)
		k = j
	}
}

func (m *MaxHeap) Replace(data int) int {
	ret := m.FindMax()

	m.array.Set(0, data)
	m.siftDown(0)

	return ret
}

func NewMaxHeap2(arr []int) *MaxHeap {

	heap := MaxHeap{}
	heap.array = list.NewArray2(arr)
	for i := heap.parent(len(arr) - 1); i > 0; i-- {
		heap.siftDown(i)
	}
	return &heap
}
