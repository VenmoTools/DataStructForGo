package unionfind

type Find5 struct {
	parent []int
	rank   []int //表示以i为根的树的高度
}

func NewUnionFind5(size int) *Find5 {
	u := Find5{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := 0; i < len(u.parent); i++ {
		u.parent[i] = i

		u.rank[i] = 1
	}
	return &u
}

// 查找p对应的编号
func (u *Find5) find(p int) int {
	if p < 0 && p >= len(u.parent) {
		panic("index error")
	}
	// 循环查找直到p所对应的根节点
	for p != u.parent[p] {
		// p的新的父节点为父节点的父节点
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

func (u *Find5) GetSize() int {
	return len(u.parent)
}

func (u *Find5) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *Find5) UnionElement(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}
	// 将rank低的集合合并到rank高的集合中
	if u.rank[pRoot] < u.rank[qRoot] {
		// 合并两个根
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		// 合并两个根
		u.parent[qRoot] = pRoot
		u.rank[qRoot] += 1
	}
}
