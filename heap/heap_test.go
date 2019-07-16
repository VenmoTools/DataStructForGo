package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap_Add(t *testing.T) {

	h := NewMaxHeap(5)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		h.Add(rand.Intn(10))
	}

	for i := 0; i < h.Size(); i++ {
		t.Log(h.ExtractMax())
	}

}
