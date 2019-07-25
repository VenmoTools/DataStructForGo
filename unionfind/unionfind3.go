package unionfind

type Find3 struct {
	parent []int
	jz     []int //表示以i为根集合中元素的个数
}

func NewUnionFind3(size int) *Find3 {
	u := Find3{
		parent: make([]int, size),
		jz:     make([]int, size),
	}
	for i := 0; i < len(u.parent); i++ {
		u.parent[i] = i

		u.jz[i] = 1
	}
	return &u
}

// 查找p对应的编号
func (u *Find3) find(p int) int {
	if p < 0 && p >= len(u.parent) {
		panic("index error")
	}
	// 循环查找直到p所对应的根节点
	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

func (u *Find3) GetSize() int {
	return len(u.parent)
}

func (u *Find3) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *Find3) UnionElement(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}
	// 如果p为根集合元素的个数小于q为根集合元素的个数
	if u.jz[pRoot] < u.jz[qRoot] {
		// 合并两个根
		u.parent[pRoot] = qRoot
		// 合并两个根的节点数量
		u.jz[pRoot] += u.jz[qRoot]
	} else {
		// 合并两个根
		u.parent[qRoot] = pRoot
		u.jz[qRoot] += u.jz[pRoot]
	}
}
