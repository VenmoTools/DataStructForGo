package list

import "testing"

func TestArray_Add(t *testing.T) {

	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}

}

func TestArray_Contains(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}

	if !arr.Contains(5) {
		t.FailNow()
	}
}

func TestArray_Find(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}

	index := arr.Find(5)
	t.Log(index)
}

func TestArray_Get(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}
	data := arr.Get(1)

	if data != 1 {
		t.FailNow()
	}

}

func TestArray_Insert(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}
	arr.Insert(10, 1)

	data := arr.Get(1)

	if data != 10 {
		t.FailNow()
	}

}

func TestArray_Remove(t *testing.T) {

	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}
	t.Log(arr)

	arr.Remove(1)

	data := arr.Get(1)

	if data == 1 {
		t.Log(arr)
		t.FailNow()
	}

}
func TestArray_RemoveFirst(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}
	t.Log(arr)

	arr.RemoveFirst()

	data := arr.Get(0)

	if data == 0 {
		t.Log(arr)
		t.FailNow()
	}
}
func TestArray_RemoveElement(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}
	t.Log(arr)

	arr.RemoveElement(4)

	data := arr.Get(3)
	t.Log(arr)

	if data == 4 {
		t.Log(arr)
		t.FailNow()
	}
}
func TestArray_RemoveLast(t *testing.T) {
	arr := NewArray(5)

	for i := 0; i < 5; i++ {
		arr.Add(i)
	}
	t.Log(arr)
	t.Log(arr.Size())

	arr.RemoveLast()
	t.Log(arr.Size())
	data := arr.Get(3)
	t.Log(arr)

	if data == 4 {
		t.Log(arr)
		t.FailNow()
	}
}
