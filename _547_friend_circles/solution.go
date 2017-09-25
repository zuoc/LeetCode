// 并查集
package _547_friend_circles

type UnionFind struct {
	parents []int
	rank    []int
	count   int
}

func (uf *UnionFind) init() {
	for i := 0; i < len(uf.parents); i++ {
		uf.parents[i] = i
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.parents[x] != x {
		uf.parents[x] = uf.find(uf.parents[x]) // 路径压缩
	}
	return uf.parents[x]
}

func (uf *UnionFind) union(x, y int) {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return
	}

	// rank 合并
	if uf.rank[px] < uf.rank[py] {
		uf.parents[px] = py
	} else {
		uf.parents[py] = px
		if uf.rank[px] == uf.rank[py] {
			uf.rank[px]++
		}
	}

	// 合并后集合数量减少
	uf.count--
}

func newUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parents: make([]int, n, n),
		rank:    make([]int, n, n),
		count:   n}
	uf.init()
	return uf
}

func findCircleNum(M [][]int) int {
	uf := newUnionFind(len(M))

	var n int = len(M)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}

	return uf.count
}
