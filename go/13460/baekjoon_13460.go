// https://www.acmicpc.net/problem/13460

package main

import "fmt"

var res int

type location struct {
	x int
	y int
}

func dfs(mat [][]byte, depth int) {
	var r, b location
	for i := range mat {
		for j := range mat[i] {
			fmt.Printf("%c", mat[i][j])
			if mat[i][j] == 'R' {
				r.x = j
				r.y = i
			}
			if mat[i][j] == 'B' {
				b.x = j
				b.y = i
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	if mat[r.y][r.x+1] == 'O' && mat[b.y][b.x+1] != 'O' {
		if depth+1 < res {
			res = depth + 1
			return
		}
	} else if mat[r.y][r.x-1] == 'O' && mat[b.y][b.x-1] != 'O' {
		if depth+1 < res {
			res = depth + 1
			return
		}
		return
	} else if mat[r.y+1][r.x] == 'O' && mat[b.y+1][b.x] != 'O' {
		if depth+1 < res {
			res = depth + 1
			return
		}
		return
	} else if mat[r.y-1][r.x] == 'O' && mat[b.y-1][b.x] != 'O' {
		if depth+1 < res {
			res = depth + 1
			return
		}
		return
	}
	for i := 0; i < 4; i++ {
		_mat := move(mat, i)
	}
}

func main() {
	var n, m int
	// var red, blue location
	res = int(^(uint(0)) >> 1)
	fmt.Scan(&n)
	fmt.Scan(&m)
	var mat = make([][]byte, n, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]byte, m, m)
	}

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		for j := 0; j < m; j++ {
			mat[i][j] = tmp[j]
			// if tmp[j] == 'R' {
			// 	red.x = j
			// 	red.y = i
			// }
			// if tmp[j] == 'B' {
			// 	blue.x = j
			// 	blue.y = i
			// }
		}
	}
	dfs(mat, 0)
	fmt.Printf("%d\n", res)
}
