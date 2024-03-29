package AVL

import (
	"datastruct/table/list"
	"fmt"
	"math"
)

type Node struct {
	data   int
	left   *Node
	right  *Node
	height int
}

type Tree struct {
	root *Node
	size int
}

func NewNode(data int) *Node {
	return &Node{data: data, height: 1}
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) Add(data int) {
	t.root = t.add(t.root, data)
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getBalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}

	return getHeight(node.left) - getHeight(node.right)
}

func rightRotate(node *Node) *Node {
	/*
	*  Before
	*         node
	*        /   \
	*       T2    z
	*      /  \
	*     y    x
	*    /
	*   g
	 */

	t2 := node.left
	x := t2.left

	t2.right = node
	node.left = x

	/*
	 * After
	 *       T2
	 *      /  \
	 *     y   node
	 *    /    / \
	 *   g    x   z
	 */
	node.height = newNodeHeight(node)
	t2.height = newNodeHeight(t2)
	return t2
}

func newNodeHeight(node *Node) int {
	return max(getHeight(node.left), getHeight(node.right)) + 1
}

func leftRotate(node *Node) *Node {

	/*
	* Before
	*       node
	*      /   \
	*     z    T2
	*         /  \
	*        x    y
	*              \
	*               g
	 */

	t2 := node.right
	x := t2.left

	t2.left = node
	node.right = x

	/*
	 * After
	 *
	 *          T2
	 *         /  \
	 *      node   y
	 *       / \    \
	 *      z   x    g
	 */
	node.height = newNodeHeight(node)
	t2.height = newNodeHeight(t2)
	return t2
}

func (t *Tree) add(node *Node, data int) *Node {
	if node == nil {
		t.size++
		return NewNode(data)
	}

	if data > node.data {
		node.right = t.add(node.right, data)
	} else if data < node.data {
		node.left = t.add(node.left, data)
	} else {
		node.data = data
	}

	node.height = newNodeHeight(node)

	nodeBalance := getBalanceFactor(node)
	// 当前平衡因子大于1并且左子树的平衡因子大于0
	/*
	 * Left
	 *     h:节点高度
	 *     b:平衡因子
	 *             T (h=4 b=2)
	 *           /   \
	 *(h=3 b=1) T2    z (h=1)
	 *         /  \
	 * (h=2)  y    x (h=1)
	 *       / \
	 * (h=1)g   h
	 */
	if nodeBalance > 1 && getBalanceFactor(node.left) >= 0 {
		return rightRotate(node)
	}

	/*
	 * Right
	 *        T (h=4 b=-2) 1-3 = -2
	 *      /   \
	 *(h=1)z    T2 (h=3 b=-1) -> 1-2
	 *         /  \
	 *  (h=1) x    y (h=2 b=-1)
	 *            / \
	 *           h   g (h=1)
	 */
	if nodeBalance < -1 && getBalanceFactor(node.right) <= 0 {
		return leftRotate(node)
	}

	/*
	*  (h=5 b=3)     T
	*              /    \
	*  (h=4 b=2) T3     T2
	*           /  \    / \
	*(h=3 b=1) T4   z  x   y
	*         /  \
	* (h=2) T5   a (h=1)
	*         \
	*          g (h=1)
	 */
	// LR
	if nodeBalance > 1 && getBalanceFactor(node.left) < 0 {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}
	// RL
	if nodeBalance < -1 && getBalanceFactor(node.right) > 0 {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func (t *Tree) FindMax() int {
	if t.size == 0 {
		return 0
	}
	return t.findMax(t.root).data
}

func (t *Tree) findMax(node *Node) *Node {
	if node.right == nil {
		return nil
	}
	return t.findMax(node.right)
}

func (t *Tree) FindMin() int {
	if t.size == 0 {
		return 0
	}
	return t.findMin(t.root).data
}

func (t *Tree) findMin(node *Node) *Node {
	if node.left == nil {
		return nil
	}
	return t.findMin(node.left)
}

func (t *Tree) Contains(data int) bool {
	return t.contains(t.root, data)
}

func (t *Tree) contains(node *Node, data int) bool {
	if node == nil {
		return false
	}

	if node.data > data {
		t.contains(node.right, data)
	} else if node.data < data {
		t.contains(node.left, data)
	}
	return true
}

func (t *Tree) RemoveMax() int {
	return t.removeMax(t.root).data
}

func (t *Tree) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		t.size--
		node.left = nil
		return leftNode
	}
	node.right = t.removeMax(node.right)
	return node
}

func (t *Tree) RemoveMin() int {
	return t.removeMin(t.root).data
}

func (t *Tree) removeMin(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		leftNode = nil
		t.size--
		return leftNode
	}

	node.left = t.removeMin(node.left)

	return node
}

func (t *Tree) Remove(data int) {
	t.remove(t.root, data)
}

func (t *Tree) remove(node *Node, data int) *Node {
	if node == nil {
		return nil
	}
	if node.data > data {
		node.right = t.remove(node.right, data)
	} else if node.data < data {
		node.left = t.remove(node.left, data)
	}
	if node.left == nil {
		rightNode := node.right
		t.size--
		node.right = nil
		return rightNode
	} else if node.right == nil {
		leftNode := node.left
		t.size--
		node.left = nil
		return leftNode
	} else {
		minNode := t.findMin(node)
		minNode.right = t.removeMin(node.right)
		minNode.left = node.left
		node.left = nil
		node.right = nil
		return minNode
	}
}

func (a *Tree) PreOrder() {
	a.order(a.root)
}

func (a *Tree) order(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	a.order(node.left)
	a.order(node.right)
}
func (a *Tree) InOrder() {
	a.inOrder(a.root)
}

func (a *Tree) inOrder(node *Node) {
	if node == nil {
		return
	}
	a.order(node.left)
	fmt.Println(node.data)
	a.order(node.right)
}

func (a *Tree) LastOrder() {
	a.lastOrder(a.root)
}

func (a *Tree) lastOrder(node *Node) {
	if node == nil {
		return
	}
	a.order(node.left)
	a.order(node.right)
	fmt.Println(node.data)
}

func (a *Tree) IsOrdered() bool {
	arr := list.NewArray(5)

	for a.Size() == 0 {

	}
	a.isOrdered(a.root, arr)

	for i := 1; i < arr.Size(); i++ {
		if arr.Get(i-1).(int) > arr.Get(i).(int) {
			return false
		}
	}
	return true
}

func (a *Tree) isOrdered(node *Node, array *list.Array) {

	if node == nil {
		return
	}
	a.isOrdered(node.left, array)
	array.Add(node.data)
	a.isOrdered(node.right, array)
}

func (a *Tree) IsBalanced() bool {
	return a.isBalanced(a.root)
}

func (a *Tree) isBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	balance := getBalanceFactor(node)
	fmt.Println("B:", balance)
	if math.Abs(float64(balance)) > 1 {
		return false
	}

	return a.isBalanced(node.left) && a.isBalanced(node.right)

}
