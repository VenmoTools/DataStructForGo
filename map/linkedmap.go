package _map

type Map interface {
	Store(key, value interface{})
	Set(key, value interface{})
	Size() int
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
}

type Node struct {
	key      interface{}
	value    interface{}
	next     *Node
	previous *Node
}

func NewNode(key, value interface{}) *Node {
	return &Node{key: key, value: value}
}

type LinkedMap struct {
	head *Node
	size int
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		head: NewNode(nil, nil),
		size: 0,
	}
}

func (l *LinkedMap) Store(key, value interface{}) {
	l.head.next = l.store(l.head.next, key, value)
}

func (l *LinkedMap) store(node *Node, key interface{}, value interface{}) *Node {
	if node == nil {
		return NewNode(key, value)
	}

	newNode := l.store(node.next, key, value)
	node.next = newNode
	newNode.previous = node
	return node
}

func (l *LinkedMap) Set(key, value interface{}) {
	node := l.head.next
	for node != nil {
		if node.key == key {
			node.value = value
		}
		node = node.next
	}
}

func (l *LinkedMap) Size() int {
	return l.size
}

func (l *LinkedMap) Contains(key interface{}) bool {
	node := l.head.next
	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}
	return false
}

func (l *LinkedMap) Get(key interface{}) interface{} {
	node := l.head.next
	for node != nil {
		if node.key == key {
			return node.value
		}
		node = node.next
	}
	return nil
}
