package binarySearchTree

import (
	"fmt"
	"testing"
)

func TestBST_Add(t *testing.T) {

	bst := NewBST()

	for i := 0; i < 50; i++ {
		bst.Add(i)
	}

	//bst.InOrder()

	fmt.Println(bst.IsOrdered())
}

func TestBST_Remove(t *testing.T) {
	bst := NewBST()

	for i := 0; i < 50; i++ {
		bst.Add(i)
	}

	bst.Remove(33)

	bst.InOrder()

	fmt.Println(bst.IsOrdered())

}
