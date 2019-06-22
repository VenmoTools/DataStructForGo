package binarySearchTree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

type BST struct {
	root *Node
	size int
}

func NewBST() *BST {
	return &BST{
		size: 0,
	}
}

func (a *BST) IsEmpty() bool {
	return a.size == 0
}

func (a *BST) Size() int {
	return a.size
}

func (a *BST) Add(data int) {
	a.root = a.add(a.root, data)
}

// 添加节点，返回添加节点后的根
func (a *BST) add(node *Node, data int) *Node {
	if node == nil {
		a.size++
		return NewNode(data)
	}
	if data < node.data {
		node.left = a.add(node.left, data)
	} else if data > node.data {
		node.right = a.add(node.right, data)
	}
	return node
}

// 判断是否含有该元素
func (a *BST) Contains(data int) bool {
	return a.contains(a.root, data)
}

func (a *BST) contains(node *Node, data int) bool {
	if node == nil {
		return false
	}
	if node.data > data {
		a.contains(node.left, data)
	} else if node.data < data {
		a.contains(node.right, data)
	}
	return true
}

func (a *BST) FindMin() int {
	if a.size == 0 {
		return 0
	}
	return a.findMin(a.root).data
}

func (a *BST) findMin(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return a.findMin(node.left)
}

func (a *BST) FindMax() int {
	if a.size == 0 {
		return 0
	}
	return a.findMax(a.root).data
}

func (a *BST) findMax(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return a.findMin(node.right)
}

func (a *BST) Remove(data int) int {
	return a.remove(a.root, data).data
}

func (a *BST) remove(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	// 只有右子树
	if data > node.data {
		// 删除右子树
		node.right = a.remove(node.right, data)
		return node
	} else if data < node.data {
		node.left = a.remove(node.left, data)
		return node
	}

	// 处理只有一个孩子的节点
	if node.left == nil {
		//如果左子树为空，只需处理右子树
		rightNode := node.right
		// 将右子树置空
		node.right = nil
		a.size--
		// 将当前的节点返回上层节点
		return rightNode
	} else if node.right == nil {
		// 如果右子树为空，只需处理左子树
		leftNode := node.left
		// 删除左子树
		node.left = nil
		a.size--
		// 将节点返回上层节点
		return leftNode
	} else {
		/*
				 5
			   /   \
			  4     7(被删除，找到8)
			 / \   / \
			2   1 6   9
					 / \
					8  11
		*/
		// 处理左右孩子都有的节点
		// 获取右子树中比该节点大的最小节点
		minNode := a.findMin(node.right)
		// 删除右子树中的最小值
		minNode.right = a.removeMin(node.right)
		// 将该节点的左子树交给后继节点
		minNode.left = node.left
		node.left = nil
		node.right = nil
		return minNode
	}
}

func (a *BST) otherRemove(node *Node, data int) *Node {
	if node == nil {
		return nil
	}

	if node.data > data {
		node.right = a.otherRemove(node.right, data)
	} else if node.data < data {
		node.left = a.otherRemove(node.left, data)
	}

	if node.left != nil && node.right != nil {
		// 获取右子树中比该节点大的最小节点
		minNode := a.findMin(node.right)
		// 把当前节点的值变为最小节点的值
		node.data = minNode.data
		// 删除
		minNode.right = a.otherRemove(node.right, data)
	} else {
		if node.left == nil {
			// 当前节点转为当前右节点
			node = node.right
		}
		if node.right == nil {
			// 当前节点转为当前左节点
			node = node.left
		}
	}
	return node
}

// 删除最小节点
func (a *BST) RemoveMin() int {
	min := a.FindMin()
	a.root = a.removeMin(a.root)
	return min
}

// 删除节点，返回删除节点后新树的根
func (a *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		a.size--
		return rightNode
	}

	node.left = a.removeMin(node.left)
	return node
}

// 删除最大节点
func (a *BST) RemoveMax() int {
	max := a.FindMax()
	// 返回删除后的节点
	a.root = a.removeMax(a.root)
	return max
}

// 删除节点，返回删除节点后新树的根
func (a *BST) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		a.size--
		return leftNode
	}

	node.right = a.removeMax(node.right)
	return node
}

func (a *BST) PreOrder() {
	a.order(a.root)
}

func (a *BST) order(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	a.order(node.left)
	a.order(node.right)
}
func (a *BST) InOrder() {
	a.inOrder(a.root)
}

func (a *BST) inOrder(node *Node) {
	if node == nil {
		return
	}
	a.order(node.left)
	fmt.Println(node.data)
	a.order(node.right)
}

func (a *BST) LastOrder() {
	a.lastOrder(a.root)
}

func (a *BST) lastOrder(node *Node) {
	if node == nil {
		return
	}
	a.order(node.left)
	a.order(node.right)
	fmt.Println(node.data)
}
