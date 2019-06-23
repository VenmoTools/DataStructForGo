package list

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

func newNode(data int) *Node {
	return &Node{data: data}
}

type LinkedList struct {
	header *Node
	size   int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{header: newNode(0)}
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Add(data int) {
	l.header.next = l.add(l.header.next, data)
}

func (l *LinkedList) add(node *Node, data int) *Node {
	if node == nil {
		l.size++
		return newNode(data)
	}
	node.next = l.add(node.next, data)
	return node
}

func (l *LinkedList) Get(index int) int {
	if index > l.size || index < 0 {
		panic("index out of range")
	}

	node := l.header.next

	for current := 0; current < index; current++ {
		node = node.next
	}

	return node.data
}

func (l *LinkedList) Insert(data int, index int) {
	if index > l.size || index < 0 {
		panic("index out of range")
	}
	node := l.header.next
	// 找到被寻找的上一个元素
	for i := 0; i < index-1; i++ {
		node = node.next
	}

	current := newNode(data)

	current.next = node.next
	node.next = current
}

func (l *LinkedList) Remove(index int) int {
	if index > l.size || index < 0 {
		panic("index out of range")
	}
	node := l.header.next
	// 找到被寻找的上一个元素
	for i := 0; i < index-1; i++ {
		node = node.next
	}

	delNode := node.next
	node.next = delNode.next
	delNode.next = nil

	return delNode.data
}

func (l *LinkedList) Index(data int) int {
	node := l.header.next
	for i := 0; node.next != nil; i++ {
		node = node.next
		if node.data == data {
			return i
		}
	}
	return -1
}

func (l *LinkedList) RemoveElement(data int) int {
	return l.Remove(l.Index(data))
}

func (l *LinkedList) Update(index int, data int) {
	if index > l.size || index < 0 {
		panic("index out of range")
	}
	node := l.header.next
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.data = data
}

// 链表反转
func (l *LinkedList) Reverse() {

	var pre, cur, next *Node
	//    pre cur
	//next
	// a -> b-> c-> d
	//
	next = l.header.next

	for next != nil {
		// pre<nil>
		//next cur
		// a->  b-> c-> d
		cur = next.next
		// pre<nil>
		// ^
		// |next cur
		// a      b-> c-> d
		next.next = pre
		//<nil>
		// ^
		// | next pre   next
		// a         b -> c-> d
		pre = next
		// <nil>
		// ^
		// | pre cur|next
		// a       b ->    c-> d
		next = cur

	}

	l.header.next = pre
}

// 链表反转2
func (l *LinkedList) Reverse2() {
	//  cur next
	//  a-> b->c->d
	var next, cur *Node
	//       cur
	// h->a-> b->c->d
	cur = l.header.next.next
	//        cur
	// h->nil a->b->c->d
	l.header.next.next = nil

	for cur != nil {
		//        cur next
		// h->nil a->  b->c->d
		next = cur.next
		//                   next
		//  h-> nil <-a(cur)   b->c->d
		cur.next = l.header.next
		//       cur        next
		//  h->  a-> nil    b->c->d
		l.header.next = cur
		//               cur|next
		//  h-> a-> nil    b->   c->d
		cur = next

	}
}

func (l *LinkedList) String() string {
	b := strings.Builder{}
	b.WriteString("[")
	for node := l.header.next; node != nil; node = node.next {
		b.WriteString(fmt.Sprintf("%d", node.data))
		if node.next != nil {
			b.WriteString("->")
		}
	}
	b.WriteString("]")
	return b.String()
}
