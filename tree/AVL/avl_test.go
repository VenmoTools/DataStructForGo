package AVL

import "testing"

func TestTree_Add(t *testing.T) {
	tree := New()

	for i := 0; i < 10; i++ {
		tree.Add(i)
	}

	t.Log(tree.Size())
}

func TestTree_Contains(t *testing.T) {
	tree := New()

	for i := 0; i < 5; i++ {
		tree.Add(i)
	}

	if !tree.Contains(4) {
		t.FailNow()
	}
}

func TestTree_InOrder(t *testing.T) {

}

func TestTree_LastOrder(t *testing.T) {

}
func TestTree_PreOrder(t *testing.T) {

}

func TestTree_RemoveMax(t *testing.T) {

}

func TestTree_FindMax(t *testing.T) {

}
func TestTree_FindMin(t *testing.T) {

}

func TestTree_Remove(t *testing.T) {

}

func TestTree_RemoveMin(t *testing.T) {

}
