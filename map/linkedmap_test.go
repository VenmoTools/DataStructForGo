package _map

import "testing"

func TestLinkedMap_Store(t *testing.T) {
	maps := NewLinkedMap()
	for i := 0; i < 10; i++ {
		maps.Store(i, i*2)
	}
}

func TestLinkedMap_Contains(t *testing.T) {
	maps := NewLinkedMap()
	for i := 0; i < 10; i++ {
		maps.Store(i, i*2)
	}
	if !maps.Contains(5) {
		t.FailNow()
	}
}

func TestLinkedMap_Get(t *testing.T) {
	maps := NewLinkedMap()
	for i := 0; i < 10; i++ {
		maps.Store(i, i*2)
	}
	if maps.Get(5).(int) != 10 {
		t.FailNow()
	}
}

func TestLinkedMap_Set(t *testing.T) {
	maps := NewLinkedMap()
	for i := 0; i < 10; i++ {
		maps.Store(i, i*2)
	}
	maps.Set(5, 100)
	if maps.Get(5).(int) != 100 {
		t.FailNow()
	}
}
