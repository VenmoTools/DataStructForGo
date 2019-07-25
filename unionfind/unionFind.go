package unionfind

type UF interface {
	GetSize() int
	IsConnected(p, q int) bool
	UnionElement(p, q int)
}

type UnionFind struct {
	data []int
}

func NewUnionFind(size int) *UnionFind {
	u := UnionFind{
		data: make([]int, size),
	}
	for i := 0; i < len(u.data); i++ {
		u.data[i] = i
	}
	return &u
}

func (u *UnionFind) GetSize() int {
	return len(u.data)
}

func (u *UnionFind) find(index int) int {
	if index < 0 && index >= len(u.data) {
		panic("index error")
	}
	return u.data[index]
}

func (u *UnionFind) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind) UnionElement(p, q int) {
	if u.IsConnected(p, q) {
		return
	}
	pid := u.find(p)
	qid := u.find(q)

	for i := 0; i < u.GetSize(); i++ {
		if u.data[i] == pid {
			u.data[i] = qid
		}
	}
}
