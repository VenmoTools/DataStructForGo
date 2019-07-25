package unionfind

type Find struct {
	parent []int
}

func NewUnionFind2(size int) *Find {
	u := Find{
		parent: make([]int, size),
	}
	for i := 0; i < len(u.parent); i++ {
		u.parent[i] = i
	}
	return &u
}

// 查找p对应的编号
func (u *Find) find(p int) int {
	if p < 0 && p >= len(u.parent) {
		panic("index error")
	}
	// 循环查找直到p所对应的根节点
	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

func (u *Find) GetSize() int {
	return len(u.parent)
}

func (u *Find) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *Find) UnionElement(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}
	u.parent[pRoot] = qRoot
}
