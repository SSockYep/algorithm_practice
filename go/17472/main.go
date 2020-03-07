// https://www.acmicpc.net/problem/17472

package main

import "fmt"

type coord struct {
	r int
	c int
}
type node struct {
	node int
	key  int
}
type island struct {
	pos []coord
}

const maxInt = int(^uint(0) >> 1)

func isEqual(a coord, b coord) bool {
	if a.r == b.r && a.c == b.c {
		return true
	}
	return false
}

func spanLands(imap [][]int, r int, c int, re *island) {
	if imap[r][c] == 1 {
		imap[r][c] = 0
		re.pos = append(re.pos, coord{r, c})f
		if r+1 < len(imap) {
			spanLands(imap, r+1, c, re)
		}
		if c+1 < len(imap[r]) {
			spanLands(imap, r, c+1, re)
		}
		if r-1 >= 0 {
			spanLands(imap, r-1, c, re)
		}
		if c-1 >= 0 {
			spanLands(imap, r, c-1, re)
		}
	}
}

func findIslands(imap [][]int) []island {
	m := make([][]int, len(imap))
	ret := make([]island, 0)
	for i := 0; i < len(imap); i++ {
		m[i] = append([]int{}, imap[i]...)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 1 {
				var tmp island
				spanLands(m, i, j, &tmp)
				ret = append(ret, tmp)

			}
		}
	}
	return ret
}

func findBridge(s island, e island, m [][]int) int {
	ret := int(^uint(0) >> 1)
	for _, sv := range s.pos {
		up := coord{sv.r - 1, sv.c}
		for up.r >= 0 && m[up.r][up.c] == 0 {
			up.r--
		}
		down := coord{sv.r + 1, sv.c}
		for down.r < len(m) && m[down.r][down.c] == 0 {
			down.r++
		}
		left := coord{sv.r, sv.c - 1}
		for left.c >= 0 && m[left.r][left.c] == 0 {
			left.c--
		}
		right := coord{sv.r, sv.c + 1}
		for right.c < len(m[right.r]) && m[right.r][right.c] == 0 {
			right.c++
		}
		for _, ev := range e.pos {
			tmp := maxInt
			flag := false
			if isEqual(ev, up) || isEqual(ev, down) {
				tmp = sv.r - ev.r
				flag = true
			}
			if isEqual(ev, left) || isEqual(ev, right) {
				tmp = sv.c - ev.c
				flag = true
			}
			if tmp < 0 {
				tmp = -tmp
			}
			if flag {
				tmp--
				if tmp > 1 && tmp < ret {
					ret = tmp
				}
			}
		}
	}
	return ret
}

func makeGraph(islands []island, islandMap [][]int) [][]int {
	g := make([][]int, len(islands))
	for i := range g {
		g[i] = make([]int, len(islands))
	}
	m := make([][]int, len(islandMap))
	for i := 0; i < len(islandMap); i++ {
		m[i] = append([]int{}, islandMap[i]...)
	}
	for i := range islands {
		for j := i; j < len(islands); j++ {
			if i != j {
				g[i][j] = findBridge(islands[i], islands[j], m)
				g[j][i] = g[i][j]
			}
		}
	}
	return g
}

func minHeap(heap *[]node) {
	i := 0
	(*heap)[0] = (*heap)[len(*heap)-1]
	(*heap) = append([]node{}, (*heap)[:len(*heap)-1]...)
	for i < len(*heap) {
		min := (i+1)*2 - 1
		if min >= len(*heap) {
			break
		}
		if min+1 < len(*heap) && (*heap)[min].key > (*heap)[min+1].key {
			min++
		}
		if (*heap)[i].key > (*heap)[min].key {
			(*heap)[i], (*heap)[min] = (*heap)[min], (*heap)[i]
			i = min
		} else {
			break
		}
	}
}

func heapify(heap *[]node, idx int) {
	for idx > 0 {
		if (*heap)[idx].key < (*heap)[idx/2].key {
			(*heap)[idx], (*heap)[idx/2] = (*heap)[idx/2], (*heap)[idx]
			idx /= 2
		} else {
			break
		}
	}
}

func prim(g [][]int) int {
	keys := make([]int, len(g))
	p := make([]int, len(g))
	queue := make([]node, len(g))
	ret := 0
	for i := range g {
		if i == 0 {
			keys[i] = 0
		} else {
			keys[i] = maxInt
		}
		p[i] = maxInt
		queue[i].node = i
		queue[i].key = maxInt
	}
	for len(queue) > 0 {
		var u int
		u = queue[0].node
		minHeap(&queue)

		for v, weight := range g[u] {
			if weight != maxInt && v != u {
				inQ := -1
				for i, vv := range queue {
					if vv.node == v {
						inQ = i
					}
				}
				if inQ != -1 && weight < keys[v] {
					keys[v] = weight
					queue[inQ].key = weight
					heapify(&queue, inQ)
					p[v] = u
				}
			}
		}
	}
	for _, v := range keys {
		if v == maxInt {
			return -1
		}
		ret += v
	}
	return ret
}

func main() {
	var n, m int
	var islandMap [][]int
	var islands []island
	var graph [][]int
	fmt.Scan(&n)
	fmt.Scan(&m)
	islandMap = make([][]int, n)
	for i := 0; i < n; i++ {
		islandMap[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&islandMap[i][j])
		}
	}
	islands = findIslands(islandMap)
	graph = makeGraph(islands, islandMap)

	result := prim(graph)
	fmt.Printf("%d\n", result)

}
