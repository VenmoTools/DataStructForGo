package list

import "testing"

func TestLinkedList_Add(t *testing.T) {
	lists := NewLinkedList()
	for i := 0; i < 10; i++ {
		lists.Add(i)
	}

	for i := 0; i < lists.Size(); i++ {
		t.Log(lists.Get(i))
	}

	lists.Reverse2()
	t.Log(lists)
}
