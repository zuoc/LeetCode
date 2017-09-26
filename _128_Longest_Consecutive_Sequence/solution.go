// 备忘录 + 并查集
package _128_Longest_Consecutive_Sequence

type unionFind struct {
	seq []int

	size []int

	maxLen int
}

func (uf *unionFind) find(x int) int {
	if uf.seq[x] != x {
		uf.seq[x] = uf.find(uf.seq[x])
	}
	return uf.seq[x]
}

func (uf *unionFind) union(x, y int) {
	sx, sy := uf.find(x), uf.find(y)
	if sx == sy {
		return
	}

	if uf.size[sx] > uf.size[sy] {
		uf.seq[sy] = sx
		uf.size[sx] += uf.size[sy]
		if uf.size[sx] > uf.maxLen {
			uf.maxLen = uf.size[sx]
		}
	} else {
		uf.seq[sx] = sy
		uf.size[sy] += uf.size[sx]
		if uf.size[sy] > uf.maxLen {
			uf.maxLen = uf.size[sy]
		}
	}

}

func newUnionFind(n int) *unionFind {
	seq := make([]int, n, n)
	siez := make([]int, n, n)
	for i := 0; i < len(seq); i++ {
		seq[i] = i
		siez[i] = 1
	}

	return &unionFind{
		seq:    seq,
		size:   siez,
		maxLen: 1,
	}
}

func longestConsecutive(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	uf := newUnionFind(n)

	v := make(map[int]int, n)

	for i := 0; i < n; i++ {
		num := nums[i]
		if _, ok := v[num]; !ok {
			if vv, ok := v[num-1]; ok {
				uf.union(vv, i)
			}
			if vv, ok := v[num+1]; ok {
				uf.union(vv, i)
			}

			v[num] = i
		}
	}

	return uf.maxLen
}
