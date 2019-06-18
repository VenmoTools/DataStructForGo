package list

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
	if a.data == nil {
		a.data = make([]interface{}, 10)
	}
	if len(a.data) == a.size-1 {
		a.resize()
	}
	a.data[a.size] = data
	a.size++
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
